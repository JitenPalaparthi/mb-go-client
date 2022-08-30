package kafka

import (
	"context"
	"errors"
	"log"

	"github.com/segmentio/kafka-go"

	"github.com/JitenPalaparthi/mb-go-client/spec"
	"github.com/JitenPalaparthi/mb-go-client/spec/common"
)

type MessageBroker[T [][]byte] struct {
	Conn   []string // This is an array for kafka
	Config any
	Err    error
	Logger *log.Logger
}

func New[T [][]byte](conn []string, config any, Logger *log.Logger) (mb *MessageBroker[T], err error) {
	if len(conn) < 1 {
		return nil, errors.New("no brokers")
	}
	return &MessageBroker[T]{Conn: conn, Config: config, Logger: Logger}, nil
}

func (mb *MessageBroker[T]) Publish(ctx context.Context, message *common.Message[T]) spec.IMessage[T] {
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: mb.Conn,
		Topic:   message.Subject,
		// assign the logger to the writer
		Logger: mb.Logger,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key: message.Data[0],
		// create an arbitrary message payload for the value
		Value: message.Data[1],
	})
	if err != nil {
		mb.Err = err
		return mb
	}
	return mb
}
func (mb *MessageBroker[T]) Subscribe(ctx context.Context, message *common.Message[T], f func(data T)) spec.IMessage[T] {
	// intialize the writer with the broker addresses, and the topic
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: mb.Conn,
		Topic:   message.Subject,
		// assign the logger to the writer
		Logger: mb.Logger,
	})

	msg, err := r.ReadMessage(ctx)
	if err != nil {
		mb.Err = err
		return mb
	}
	f([][]byte{msg.Key, msg.Value})

	return mb
}
func (mb *MessageBroker[T]) SubscribeSync(ctx context.Context, message *common.Message[T], f func(data T)) spec.IMessage[T] {
	// intialize the writer with the broker addresses, and the topic
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: mb.Conn,
		Topic:   message.Subject,
		// assign the logger to the writer
		Logger: mb.Logger,
	})

	msg, err := r.ReadMessage(ctx)
	if err != nil {
		mb.Err = err
		return mb
	}
	f([][]byte{msg.Key, msg.Value})

	return mb
}

func (mb *MessageBroker[T]) Error() error {
	return mb.Err
}
