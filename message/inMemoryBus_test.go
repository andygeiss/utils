package message_test

import (
	"github.com/andygeiss/utils/assert"
	"github.com/andygeiss/utils/message"
	"testing"
	"time"
)

func TestInMemoryBus(t *testing.T) {

	bus := message.DefaultBus

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

	bus.Subscribe("foo", consumer1Ch)
	bus.Subscribe("foo", consumer2Ch)

	// producer
	bus.Publish("foo", "bar")

	time.Sleep(time.Second * 1)

	assert.That("data1 should be bar", t, data1, "bar")
	assert.That("data2 should be bar", t, data2, "bar")
}
