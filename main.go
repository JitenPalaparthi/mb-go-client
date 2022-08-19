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

	var ifc1 spec.Messager[[]byte]
	//var ifc2 spec.Messager[[][]byte]
	mb1 := new(nats.MessageBroker[[]byte])
	mb1.Conn = "nats://localhost:4222"
	//mb1.Subject="demos.demo1"
	//mb1.Data=[]byte("Hello Muruga")
	ifc1 = mb1

	msg1 := new(spec.Message[[]byte])
	msg1.Subject = "demo.demo1"
	msg1.Data = []byte("Hello Muruga")

	msg2 := new(spec.Message[[]byte])
	msg2.Subject = "demo.demo1"

	msg3 := new(spec.Message[[]byte])
	msg3.Subject = "demo.demo2"
	msg3.Data = []byte("Hello Muruga -- again")

	msg4 := new(spec.Message[[]byte])
	msg4.Subject = "demo.demo2"

	go func() {
		for {
			select {
			case err := <-ifc1.OnErr():
				fmt.Println("here is error", err)
				// default:
				// 	fmt.Println("------------------>")
			}
		}
	}()

	//go ifc1.OnErr()
	go ifc1.Publish(context.TODO(), msg1).Subscribe(context.TODO(), msg2, func(data []byte) {
		fmt.Println(string(data) + " You are the world-1")
	}).Publish(context.TODO(), msg3).Subscribe(context.TODO(), msg4, func(data []byte) {
		fmt.Println(string(data) + " You are the world-2")
	})

	// mb2 := new(kafka.MessageBroker[[][]byte])
	// mb2.Conn = []string{"localhost:29092"}
	// ifc2 = mb2

	// go ifc2.Publish(context.TODO(), "demo.demo3", [][]byte{[]byte("name"), []byte("Muruga be with me")}, chErr)
	// go ifc2.Subscribe(context.TODO(), "demo.demo3", func(data [][]byte) {
	// 	fmt.Println("Received from-->", data)
	// }, chErr)

	runtime.Goexit()
}
