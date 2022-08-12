package broker

// Publisher is to publish a message on a topic/subject
type Publisher interface {
	Publish() error
}

// Subscriber is used to subscribe a message on a subject/topic
type Subscriber interface {
	Subscribe(func()) error
}

// Processor is used to process a message
// that means a message is subscribed from one topic/subject and published to another topic/subject
type Procesor interface {
	Process() error
}
