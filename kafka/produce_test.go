package kafka

import (
	"fmt"
	"testing"
)

func TestProduce(t *testing.T) {
	Segmentio([]Message{{Key: "Segmentio", Value: "hello world"}})
	Confluent([]Message{{Key: "Confluent", Value: "hello world"}})
	Sarama([]Message{{Key: "Sarama", Value: "hello world"}})
	Goka([]Message{{Key: "Goka", Value: "hello world"}})
}

func BenchmarkProduceSegmentio(b *testing.B) {
	count := b.N

	messages := make([]Message, count)

	for i := 0; i < count; i++ {
		messages[i] = Message{Key: "Segmentio", Value: fmt.Sprint(i)}
	}

	Segmentio(messages)
}

func BenchmarkProduceConfluent(b *testing.B) {
	count := b.N

	messages := make([]Message, count)

	for i := 0; i < count; i++ {
		messages[i] = Message{Key: "Confluent", Value: fmt.Sprint(i)}
	}

	Confluent(messages)
}

func BenchmarkProduceSarama(b *testing.B) {
	count := b.N

	messages := make([]Message, count)

	for i := 0; i < count; i++ {
		messages[i] = Message{Key: "Sarama", Value: fmt.Sprint(i)}
	}

	Sarama(messages)
}

func BenchmarkProduceGoka(b *testing.B) {
	count := b.N

	messages := make([]Message, count)

	for i := 0; i < count; i++ {
		messages[i] = Message{Key: "Goka", Value: fmt.Sprint(i)}
	}

	Goka(messages)
}
