package nats

import (
	"github.com/nats-io/nats.go"
)

type MessageBroker struct {
	Conn    string
	Configs any
}

type Message struct {
	Subject string
	Data    []byte
}

func (pm *MessageBroker) Publish(msg Message) error {
	nc, err := nats.Connect(pm.Conn)
	if err != nil {
		return err
	}
	defer nc.Close()

	if err := nc.Publish(msg.Subject, msg.Data); err != nil {
		return err
	}
	return nil
}

func (mb *MessageBroker) Subscribe(msg Message, f func(data []byte)) error {
	nc, err := nats.Connect(mb.Conn)
	if err != nil {
		return err
	}
	_, err = nc.Subscribe(msg.Subject, func(m *nats.Msg) {
		msg.Data = m.Data
		f(msg.Data)
	})
	return err
}

func (mb *MessageBroker) Process(f func(data []byte)) error {
	nc, err := nats.Connect(mb.Conn)
	if err != nil {
		return err
	}
	_, err = nc.Subscribe(mb.Subject, func(m *nats.Msg) {
		mb.Data = m.Data
		f(mb.Data)
	})
	return err
}
