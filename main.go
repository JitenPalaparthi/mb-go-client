package main

import (
	"fmt"
	"messagebroker/implementers/nats"
	"runtime"
)

func main() {
	mb := new(nats.MessageBroker)
	mb.Conn = "nats://localhost:4222"
	mb.Subject = "demos.demo"
	mb.Data = []byte("Hello Muruga")

	if err := mb.Process(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Published")

	f := func(data []byte) {
		if data != nil {
			fmt.Println(string(data))
		} else {
			fmt.Println("nil it is")
		}
	}
	go func() {
		mb.Subscribe(f)
	}()
	runtime.Goexit()
}
