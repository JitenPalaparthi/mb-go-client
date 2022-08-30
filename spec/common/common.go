package common

type Type interface {
	[]byte | [][]byte
}

type Error interface {
	Error() error
}

type Message[T Type] struct {
	Subject string
	Data    T
	Config  any
}
