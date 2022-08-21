package spec

import "context"

type Message[T Type] struct {
	Subject string
	Data    T
	Config  any
}

// Messager is an interfer that implements two other interfaces which are Publisher and Subscriber
type Messager[T Type] interface {
	Publisher[T]
	Subscriber[T]
	Error
}

type Type interface {
	[]byte | [][]byte
}

// Publisher is to publish a message on a topic/subject
type Publisher[T Type] interface {
	Publish(ctx context.Context, message *Message[T]) Messager[T]
}

// Subscriber is used to subscribe a message on a subject/topic
type Subscriber[T Type] interface {
	Subscribe(ctx context.Context, message *Message[T], f func(data T)) Messager[T]
}

type Error interface {
	Error() error
}
