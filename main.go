package main

import (
	"fmt"
	"runtime"

	"github.com/JitenPalaparthi/mb-go-client/impl"
	"github.com/JitenPalaparthi/mb-go-client/spec"
)

func main() {
	mb := new(impl.MessageBroker)
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

	var ifc spec.Messager

	mb1 := new(impl.MessageBroker)
	mb1.Conn = "nats://localhost:4222"
	//mb1.Subject="demos.demo1"
	//mb1.Data=[]byte("Hello Muruga")
	ifc = mb1

	go ifc.Publish("demo.demo1", []byte("Hello Muruga")).Subscribe("demo.demo1", func(data []byte) {
		fmt.Println(string(data) + " You are the world-1")
	}).Publish("demo.demo2", []byte("Hello Muruga -- again")).Subscribe("demo.demo2", func(data []byte) {
		fmt.Println(string(data) + " You are the world-2")
	})
	runtime.Goexit()
}
