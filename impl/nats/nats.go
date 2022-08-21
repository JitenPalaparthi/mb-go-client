package nats

import (
	"context"
	"errors"

	"github.com/JitenPalaparthi/mb-go-client/spec"
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

func (mb *MessageBroker[T]) Publish(ctx context.Context, message *spec.Message[T]) spec.Messager[T] {
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
		mb.Err = err
		return mb
	}
	return mb
}

func (mb *MessageBroker[T]) Subscribe(ctx context.Context, message *spec.Message[T], f func(data T)) spec.Messager[T] {
	if mb.Err != nil {
		return mb
	}
	nc, err := nats.Connect(mb.Conn)
	if err != nil {
		mb.Err = err
		return mb
	}
	_, err = nc.Subscribe(message.Subject, func(m *nats.Msg) {
		if err != nil {
			mb.Err = err
			return
		}
		f(m.Data)
		m.Sub.Unsubscribe()
	})
	return mb
}

func (mb *MessageBroker[T]) Error() error {
	return mb.Err
}
