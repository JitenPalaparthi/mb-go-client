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
		err = messager.Subscribe(context.TODO(), &common.Message[[]byte]{Subject: "demo.demo1"}, func(data []byte) {
			fmt.Println("subscribe:", string(data))
		}).Error()
		if err != nil {
			fmt.Println(err)
		}
	}

	runtime.Goexit()
}
