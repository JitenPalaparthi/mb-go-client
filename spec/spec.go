package spec

import "context"

type Message[T DataType] struct {
	Subject string
	Data    T
	Config  interface{}
	//ChErr   chan error
}

// Messager is an interfer that implements two other interfaces which are Publisher and Subscriber
type Messager[T DataType] interface {
	Publisher[T]
	Subscriber[T]
	Error
}

type DataType interface {
	[]byte | [][]byte
}

// Publisher is to publish a message on a topic/subject
type Publisher[T DataType] interface {
	Publish(ctx context.Context, message *Message[T]) Messager[T]
}

// Subscriber is used to subscribe a message on a subject/topic
type Subscriber[T DataType] interface {
	Subscribe(ctx context.Context, message *Message[T], f func(data T)) Messager[T]
}

type Error interface {
	OnErr() <-chan error
}
