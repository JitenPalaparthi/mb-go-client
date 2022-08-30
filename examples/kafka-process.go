package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/JitenPalaparthi/mb-go-client/impl/kafka"
	"github.com/JitenPalaparthi/mb-go-client/spec"
	"github.com/JitenPalaparthi/mb-go-client/spec/common"
)

func main() {
	// nats example
	var (
		messager spec.IMessage[[][]byte]
		err      error
	)
	l := log.New(os.Stdout, "kafka writer: ", 0)

	//kafka.New[[][]byte]()
	if messager, err = kafka.New[[][]byte]([]string{"localhost:29092"}, nil, l); err != nil {
		fmt.Println(err)
	} else {

		go func() {
			err = messager.SubscribeSync(context.TODO(), &common.Message[[][]byte]{Subject: "demo.demo2"}, func(data [][]byte) {
				fmt.Println("Subscribe:", "Key:", string(data[0]), "Value:", string(data[1]))
			}).Error()
			if err != nil {
				fmt.Println(err)
			}
		}()
		go func() {
			dta1 := [][]byte{[]byte("message"), []byte("Hello World-2")}
			err = messager.SubscribeSync(context.TODO(), &common.Message[[][]byte]{Subject: "demo.demo1"}, func(data [][]byte) {
				fmt.Println("Subscribe:", "Key:", string(data[0]), "Value:", string(data[1]))
			}).Publish(context.TODO(), &common.Message[[][]byte]{Subject: "demo.demo2", Data: dta1}).Error()
			if err != nil {
				fmt.Println(err)
			}
		}()
		// publisher
		go func() {
			dta2 := [][]byte{[]byte("message"), []byte("Hello World-1")}

			err = messager.Publish(context.TODO(), &common.Message[[][]byte]{Subject: "demo.demo1", Data: dta2}).Error()
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	runtime.Goexit()
}
