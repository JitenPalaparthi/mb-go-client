package spec

type Message struct {
	Subject string
	Data    []byte
}

// Publisher is to publish a message on a topic/subject
type Publisher interface {
	Publish(*Message) error
}

// Subscriber is used to subscribe a message on a subject/topic
type Subscriber interface {
	Subscribe(*Message, func(data []byte)) error
}

// Processor is used to process a message
// that means a message is subscribed from one topic/subject and published to another topic/subject
type Procesor interface {
	Process(*Message, func(data []byte) *Message) error
}
