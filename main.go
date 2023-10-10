package main

import (
	"time"

	"github.com/dragno99/cache-service/client"
	"github.com/dragno99/cache-service/server"
)

func main() {

	// starting server in another go routine so that it wont block our code
	go server.StartServer()

	// starting Custom User Client server in another go routine so that it wont block our code
	go server.StartUserClientServer()

	time.Sleep(time.Millisecond * 5000)

	// calling Test method to test the client
	client.Test()

	// calling Test method to test the Custom User client
	client.TestCustomUserClient()

	<-make(chan struct{})
}
