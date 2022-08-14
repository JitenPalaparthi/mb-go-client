package impl

import (
	"context"

	"github.com/JitenPalaparthi/mb-go-client/spec"
	"github.com/nats-io/nats.go"
)

type MessageBroker struct {
	Conn   any
	Config any
	Err    error
}

func (mb *MessageBroker) Publish(ctx context.Context, subject string, data []byte) spec.Messager {
	nc, err := nats.Connect(mb.Conn.(string))
	if err != nil {
		mb.Err = err
		return mb
	}
	defer nc.Close()

	if err := nc.Publish(subject, data); err != nil {
		mb.Err = err
		return mb
	}
	return mb
}

func (mb *MessageBroker) Subscribe(ctx context.Context, subject string, f func(data []byte)) spec.Messager {
	nc, err := nats.Connect(mb.Conn.(string))
	if err != nil {
		mb.Err = err
		return mb
	}
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		if err != nil {
			mb.Err = err
			return
		}
		f(m.Data)
	})
	return mb
}
