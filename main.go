package main

import (
	"fmt"
	"runtime"

	"github.com/JitenPalaparthi/mb-go-client/mbs/nats"
	"github.com/JitenPalaparthi/mb-go-client/spec"
)

func main() {
	mb := new(nats.BrokerConfig)
	mb.Conn = "nats://localhost:4222"
	//mb.Subject = "demos.demo"
	//mb.Data = []byte("Hello Muruga")

	msg := &spec.Message{Subject: "demos.demo", Data: []byte("Hello Murua")}
	if err := mb.Publish(msg); err != nil {
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
		mb.Subscribe(msg, f)
	}()
	runtime.Goexit()
}
