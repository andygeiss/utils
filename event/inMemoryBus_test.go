package event_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/event"
	"testing"
	"time"
)

func TestInMemoryBus(t *testing.T) {

	bus := event.DefaultBus

	// consumer 1
	consumer1Ch := make(chan interface{})
	data1 := ""
	go func() {
		val := <-consumer1Ch
		data1 = val.(string)
	}()

	// consumer 1
	consumer2Ch := make(chan interface{})
	data2 := ""
	go func() {
		val := <-consumer2Ch
		data2 = val.(string)
	}()

	bus.Register("foo", consumer1Ch)
	bus.Register("foo", consumer2Ch)

	// producer
	bus.Publish("foo", "bar")

	time.Sleep(time.Second * 1)

	assert.That("data1 should be bar", t, data1, "bar")
	assert.That("data2 should be bar", t, data2, "bar")
}
