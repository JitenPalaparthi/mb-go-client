package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/JitenPalaparthi/mb-go-client/impl/nats"
	"github.com/JitenPalaparthi/mb-go-client/spec"
	"github.com/JitenPalaparthi/mb-go-client/spec/common"
)

func main() {
	// nats example
	var (
		messager spec.IMessage[[]byte]
		err      error
	)

	if messager, err = nats.New[[]byte]("nats://localhost:4222", nil); err != nil {
		fmt.Println(err)
	} else {

		go func() {
			err = messager.SubscribeSync(context.TODO(), &common.Message[[]byte]{Subject: "demo.demo2"}, func(data []byte) {
				fmt.Println("Subscribe:", string(data))
			}).Error()
			if err != nil {
				fmt.Println(err)
			}
		}()
		go func() {
			err = messager.SubscribeSync(context.TODO(), &common.Message[[]byte]{Subject: "demo.demo1"}, func(data []byte) {
				fmt.Println("Subscribe:", string(data))
			}).Publish(context.TODO(), &common.Message[[]byte]{Subject: "demo.demo2", Data: []byte("Hello World! This is another message")}).Error()
			if err != nil {
				fmt.Println(err)
			}
		}()
		// publisher
		go func() {
			err = messager.Publish(context.TODO(), &common.Message[[]byte]{Subject: "demo.demo1", Data: []byte("Hello World! for publish only")}).Error()
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	runtime.Goexit()
}
