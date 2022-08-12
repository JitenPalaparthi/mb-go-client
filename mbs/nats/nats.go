package nats

import (
	spec "github.com/JitenPalaparthi/mb-go-client/spec"
	"github.com/nats-io/nats.go"
)

type BrokerConfig struct {
	Conn    string
	Configs any
}

func (bc *BrokerConfig) Publish(msg *spec.Message) error {
	nc, err := nats.Connect(bc.Conn)
	if err != nil {
		return err
	}
	defer nc.Close()

	if err := nc.Publish(msg.Subject, msg.Data); err != nil {
		return err
	}
	return nil
}

func (bc *BrokerConfig) Subscribe(msg *spec.Message, f func(data []byte)) error {
	nc, err := nats.Connect(bc.Conn)
	if err != nil {
		return err
	}
	_, err = nc.Subscribe(msg.Subject, func(m *nats.Msg) {
		msg.Data = m.Data
		f(msg.Data)
	})
	return err
}

func (bc *BrokerConfig) Process(msg *spec.Message, f func(data []byte) *spec.Message) error {
	nc, err := nats.Connect(bc.Conn)
	if err != nil {
		return err
	}
	_, err = nc.Subscribe(msg.Subject, func(m *nats.Msg) {
		msg.Data = m.Data
		pmsg := f(msg.Data)
		bc.Publish(pmsg)
	})
	return err
}
