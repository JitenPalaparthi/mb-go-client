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
	Err    error
}

func (mb *MessageBroker[T]) Publish(ctx context.Context, subject string, data T) spec.Messager[T] {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: mb.Conn.([]string),
		Topic:   subject,
		// assign the logger to the writer
		Logger: l,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key: data[0],
		// create an arbitrary message payload for the value
		Value: data[1],
	})
	if err != nil {
		mb.Err = err
		return mb
	}
	return mb
}

func (mb *MessageBroker[T]) Subscribe(ctx context.Context, subject string, f func(data T)) spec.Messager[T] {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: mb.Conn.([]string),
		Topic:   subject,
		// assign the logger to the writer
		Logger: l,
	})

	msg, err := r.ReadMessage(ctx)
	if err != nil {
		mb.Err = err
		return mb
	}
	var data [][]byte
	data[0] = msg.Key
	data[1] = msg.Value
	f(data)

	return mb
}
