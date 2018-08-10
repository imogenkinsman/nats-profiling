package main

import (
	"fmt"

	stand "github.com/nats-io/nats-streaming-server/server"
)

func main() {

	go func() {
		_, err := stand.RunServer("mystreamingserver")
		fmt.Print("I did a thing\n")

		if err != nil {
			fmt.Errorf("Failed to start NATS streaming server: %v\n", err)
		}
	}()

	select {}
}
