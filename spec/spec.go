package spec

import "context"

// Messager is an interfer that implements two other interfaces which are Publisher and Subscriber
type Messager interface {
	Publisher
	Subscriber
}

// Publisher is to publish a message on a topic/subject
type Publisher interface {
	Publish(ctx context.Context, subject string, data []byte) Messager
}

// Subscriber is used to subscribe a message on a subject/topic
type Subscriber interface {
	Subscribe(ctx context.Context, subject string, f func(data []byte)) Messager
}
