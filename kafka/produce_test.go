package kafka

import (
	"fmt"
	"testing"
)

func TestProduce(t *testing.T) {
	Segmentio(Message{Key: "Segmentio", Value: "hello world"})
	Confluent(Message{Key: "Confluent", Value: "hello world"})
	Sarama(Message{Key: "Sarama", Value: "hello world"})
	Goka(Message{Key: "Goka", Value: "hello world"})
}

func TestBulkProduce(t *testing.T) {
	count := 50

	bulk := make([]Message, count)

	for i := 0; i < count; i++ {
		bulk[i] = Message{Key: "Bulk", Value: fmt.Sprint(i)}
	}

	for _, v := range bulk {
		Segmentio(v)
		Confluent(v)
		Sarama(v)
		Goka(v)
	}
}
