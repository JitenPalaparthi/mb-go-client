package kafka

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"

	"github.com/JitenPalaparthi/mb-go-client/spec"
)

type MessageBroker[T [][]byte] struct {
	Conn   any // This is an array for kafka
	Config any
	ChErr  chan error
}

func (mb *MessageBroker[T]) Publish(ctx context.Context, message *spec.Message[T]) spec.Messager[T] {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: mb.Conn.([]string),
		Topic:   message.Subject,
		// assign the logger to the writer
		Logger: l,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key: message.Data[0],
		// create an arbitrary message payload for the value
		Value: message.Data[1],
	})
	if err != nil {
		if mb.ChErr == nil {
			mb.ChErr = make(chan error)
		}
		mb.ChErr <- err
		return mb
	}
	return mb
}

func (mb *MessageBroker[T]) Subscribe(ctx context.Context, message *spec.Message[T], f func(data T)) spec.Messager[T] {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: mb.Conn.([]string),
		Topic:   message.Subject,
		// assign the logger to the writer
		Logger: l,
	})

	msg, err := r.ReadMessage(ctx)
	if err != nil {
		if mb.ChErr == nil {
			mb.ChErr = make(chan error)
		}
		mb.ChErr <- err
		return mb
	}
	f([][]byte{msg.Key, msg.Value})

	return mb
}

func (mb *MessageBroker[T]) OnErr() <-chan error {
	if mb.ChErr == nil {
		mb.ChErr = make(chan error)
	}
	return mb.ChErr
}
