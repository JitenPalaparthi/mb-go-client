package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/JitenPalaparthi/mb-go-client/impl/nats"
	"github.com/JitenPalaparthi/mb-go-client/spec"
)

func main() {
	mb := new(nats.MessageBroker[[]byte])
	mb.Conn = "nats://localhost:4222"
	//mb.Subject = "demos.demo"
	//mb.Data = []byte("Hello Muruga")

	// msg := &impl.MessageBroker{Subject: "demos.demo", Data: []byte("Hello Murua")}
	// if err := mb.Publish(msg); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Published")

	// f := func(data []byte) {
	// 	if data != nil {
	// 		fmt.Println(string(data))
	// 	} else {
	// 		fmt.Println("nil it is")
	// 	}
	// }
	// go func() {
	// 	mb.Subscribe(msg, f)
	// }()

	var ifc spec.Messager[[]byte]

	mb1 := new(nats.MessageBroker[[]byte])
	mb1.Conn = "nats://localhost:4222"
	//mb1.Subject="demos.demo1"
	//mb1.Data=[]byte("Hello Muruga")
	ifc = mb1

	go ifc.Publish(context.TODO(), "demo.demo1", []byte("Hello Muruga")).Subscribe(context.TODO(), "demo.demo1", func(data []byte) {
		fmt.Println(string(data) + " You are the world-1")
	}).Publish(context.TODO(), "demo.demo2", []byte("Hello Muruga -- again")).Subscribe(context.TODO(), "demo.demo2", func(data []byte) {
		fmt.Println(string(data) + " You are the world-2")
	})
	runtime.Goexit()
}
