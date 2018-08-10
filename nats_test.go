package main

import (
	"fmt"
	"testing"

	stan "github.com/nats-io/go-nats-streaming"
	stand "github.com/nats-io/nats-streaming-server/server"
	"github.com/nats-io/nats-streaming-server/stores"
)

const ServerID = "nats"
const ClientID = "client"

func BenchmarkServer(b *testing.B) {
	// set up server
	opts := stand.GetDefaultOptions()
	opts.ID = ServerID
	_, err := stand.RunServerWithOpts(opts, nil)
	if err != nil {
		fmt.Errorf("Failed to start NATS streaming server: %v\n", err)
	}

	// connect and publish to server
	sc, err := stan.Connect(ServerID, ClientID)
	if err != nil {
		panic(err.Error())
	}

	ah := func(guid string, err error) {
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}

	data := []byte("data")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sc.PublishAsync("subject", data, ah)
	}

	sc.Close()
}

func BenchmarkPersistentServer(b *testing.B) {
	opts := stand.GetDefaultOptions()
	opts.StoreType = stores.TypeFile
	opts.FilestoreDir = "datastore"
	_, err := stand.RunServerWithOpts(opts, nil)
	if err != nil {
		fmt.Errorf("Failed to start NATS streaming server: %v\n", err)
	}

	// connect and publish to server
	sc, err := stan.Connect(ServerID, ClientID)
	if err != nil {
		panic(err.Error())
	}

	ah := func(guid string, err error) {
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}

	data := []byte("data")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sc.PublishAsync("subject", data, ah)
	}

	sc.Close()
}
