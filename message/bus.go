package message

// Bus ...
type Bus interface {
	Publish(topic string, data interface{})
	Subscribe(topic string, consumer chan interface{})
}
