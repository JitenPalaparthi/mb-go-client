package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/JitenPalaparthi/mb-go-client/impl/nats"
	"github.com/JitenPalaparthi/mb-go-client/spec"
)

func main() {
	// nats example
	var (
		messager spec.Messager[[]byte]
		err      error
	)

	if messager, err = nats.New[[]byte]("nats://localhost:4222", nil); err != nil {
		fmt.Println(err)
	} else {
		err = messager.Publish(context.TODO(), &spec.Message[[]byte]{Subject: "demo.demo1", Data: []byte("Hello World! for publish only")}).Error()
		if err != nil {
			fmt.Println(err)
		}
	}

	runtime.Goexit()
}
