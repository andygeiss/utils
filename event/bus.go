package event

// Bus ...
type Bus interface {
	Publish(topic string, data interface{})
	Register(topic string, consumer chan interface{})
}
