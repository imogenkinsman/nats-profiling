package main

import (
	"fmt"
	"testing"

	stan "github.com/nats-io/go-nats-streaming"
	stand "github.com/nats-io/nats-streaming-server/server"
)

const ServerID = "nats"
const ClientID = "client"
const OneMillion = 10

func BenchmarkServer(b *testing.B) {
	// set up server
	opts := stand.GetDefaultOptions()
	opts.ID = ServerID
	_, err := stand.RunServerWithOpts(opts, nil)
	if err != nil {
		fmt.Errorf("Failed to start NATS streaming server: %v\n", err)
	}

	// connect and publish to server
	println("test")
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
}

// func TestMain(m *testing.M) {
//
// }

// func BenchmarkPersistentServer(b *testing.B) {
// opts := stand.GetDefaultOptions()
// opts.StoreType = stores.TypeFile
// opts.FilestoreDir = "datastore"
// server, err := stand.RunServerWithOpts(opts, nil)
// if err != nil {
// 	fmt.Errorf("Failed to start NATS streaming server: %v\n", err)
// }
// }
