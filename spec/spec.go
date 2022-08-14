package spec

import "context"

// Messager is an interfer that implements two other interfaces which are Publisher and Subscriber
type Messager[T TYPE] interface {
	Publisher[T]
	Subscriber[T]
}

type TYPE interface {
	[]byte | [][]byte
}

// Publisher is to publish a message on a topic/subject
type Publisher[T TYPE] interface {
	Publish(ctx context.Context, subject string, data T) Messager[T]
}

// Subscriber is used to subscribe a message on a subject/topic
type Subscriber[T TYPE] interface {
	Subscribe(ctx context.Context, subject string, f func(data T)) Messager[T]
}
