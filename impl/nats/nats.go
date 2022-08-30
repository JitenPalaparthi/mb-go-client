package nats

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/JitenPalaparthi/mb-go-client/spec"
	"github.com/JitenPalaparthi/mb-go-client/spec/common"

	"github.com/nats-io/nats.go"
)

type MessageBroker[T []byte] struct {
	Conn   string
	Config any
	Err    error
}

func New[T []byte](conn string, config any) (mb *MessageBroker[T], err error) {
	if conn == "" {
		return nil, errors.New("no connections")
	}
	return &MessageBroker[T]{Conn: conn, Config: config}, nil
}

func (mb *MessageBroker[T]) Publish(ctx context.Context, message *common.Message[T]) spec.IMessage[T] {
	if message == nil {
		return nil
	}
	if mb.Err != nil {
		return mb
	}
	nc, err := nats.Connect(mb.Conn)
	if err != nil {
		mb.Err = err
		return mb
	}
	defer nc.Close()

	if err := nc.Publish(message.Subject, message.Data); err != nil {
		fmt.Println(err)
		mb.Err = err
		return mb
	}
	return mb
}

// Subscribe must be a sync operation. Becasue the way it will be called
func (mb *MessageBroker[T]) Subscribe(ctx context.Context, message *common.Message[T], f func(data T)) spec.IMessage[T] {

	if mb.Err != nil {
		return mb
	}
	nc, err := nats.Connect(mb.Conn)
	fmt.Println(err)

	if err != nil {
		mb.Err = err

		return mb
	}

	// Use the response
	//log.Printf("Reply: %s", msg.Data)
	_, err = nc.Subscribe(message.Subject, func(m *nats.Msg) {
		if err != nil {
			mb.Err = err
			return
		}
		f(m.Data)
		//	m.Sub.Unsubscribe()
	})

	return mb
}

// Subscribe must be a sync operation. Becasue the way it will be called
func (mb *MessageBroker[T]) SubscribeSync(ctx context.Context, message *common.Message[T], f func(data T)) spec.IMessage[T] {

	if mb.Err != nil {
		return mb
	}
	nc, err := nats.Connect(mb.Conn)
	fmt.Println(err)

	if err != nil {
		mb.Err = err

		return mb
	}

	sub, err := nc.SubscribeSync(message.Subject)
	if err != nil {
		mb.Err = err
		return mb
	}
	msg, err := sub.NextMsg(10 * time.Second)
	if err != nil {
		mb.Err = err
		return mb
	}
	f(msg.Data)
	return mb
}

func (mb *MessageBroker[T]) Error() error {
	return mb.Err
}
