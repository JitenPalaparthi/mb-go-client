package spec

import (
	"context"

	"github.com/JitenPalaparthi/mb-go-client/spec/common"
)

// Messager is an interfer that implements two other interfaces which are Publisher and Subscriber
type IMessage[T common.Type] interface {
	Publisher[T]
	Subscriber[T]
	SubscriberSync[T]
	common.Error
}

// Publisher is to publish a message on a topic/subject
type Publisher[T common.Type] interface {
	Publish(ctx context.Context, message *common.Message[T]) IMessage[T]
}

// Subscriber is used to subscribe a message on a subject/topic
type Subscriber[T common.Type] interface {
	Subscribe(ctx context.Context, message *common.Message[T], f func(data T)) IMessage[T]
}

// Subscriber is used to subscribe a message on a subject/topic
type SubscriberSync[T common.Type] interface {
	SubscribeSync(ctx context.Context, message *common.Message[T], f func(data T)) IMessage[T]
}
