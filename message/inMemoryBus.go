package message

type inMemoryBus struct {
	topics map[string][]chan interface{}
}

func (b *inMemoryBus) Publish(topic string, data interface{}) {
	if consumers, ok := b.topics[topic]; ok {
		for _, consumer := range consumers {
			consumer := consumer
			go func() {
				consumer <- data
			}()
		}
	}
}

func (b *inMemoryBus) Register(topic string, consumer chan interface{}) {
	consumers, ok := b.topics[topic]
	if !ok {
		consumers = make([]chan interface{}, 0)
	}
	consumers = append(consumers, consumer)
	b.topics[topic] = consumers
}

// NewInMemoryBus ...
func NewInMemoryBus() Bus {
	return &inMemoryBus{
		topics: map[string][]chan interface{}{},
	}
}

// DefaultBus ...
var DefaultBus = NewInMemoryBus()
