package nats

import (
	"context"
	"fmt"

	"github.com/JitenPalaparthi/mb-go-client/spec"
	"github.com/nats-io/nats.go"
)

type MessageBroker[T []byte] struct {
	Conn   any
	Config any
	ChErr  chan error
}

func (mb *MessageBroker[T]) Publish(ctx context.Context, message *spec.Message[T]) spec.Messager[T] {
	if message == nil {
		return nil
	}
	nc, err := nats.Connect(mb.Conn.(string))
	if err != nil {
		if mb.ChErr == nil {
			mb.ChErr = make(chan error)
		}

		mb.ChErr <- err
		fmt.Println(err)
		return mb
	}
	defer nc.Close()

	if err := nc.Publish(message.Subject, message.Data); err != nil {
		if mb.ChErr == nil {
			mb.ChErr = make(chan error)
		}
		mb.ChErr <- err
		return mb
	}
	return mb
}

func (mb *MessageBroker[T]) Subscribe(ctx context.Context, message *spec.Message[T], f func(data T)) spec.Messager[T] {
	nc, err := nats.Connect(mb.Conn.(string))
	if err != nil {
		if mb.ChErr == nil {
			mb.ChErr = make(chan error)
		}
		mb.ChErr <- err
		return mb
	}
	_, err = nc.Subscribe(message.Subject, func(m *nats.Msg) {
		if err != nil {
			if mb.ChErr == nil {
				mb.ChErr = make(chan error)
			}
			mb.ChErr <- err
			return
		}
		f(m.Data)
	})
	return mb
}

func (mb *MessageBroker[T]) OnErr() <-chan error {
	if mb.ChErr == nil {
		mb.ChErr = make(chan error)
	}
	return mb.ChErr
}
