package kafka

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	confluent "github.com/confluentinc/confluent-kafka-go/kafka"
	segmentio "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()

var host = "192.168.1.4:9092"
var partition = 0
var topic = "benchmark"

// Message global message structure for kafka
type Message struct {
	Key   string
	Value string
}

// Segmentio library
func Segmentio(message Message) {
	conn, err := segmentio.DialLeader(context.Background(), "tcp", host, topic, partition)

	if err != nil {
		log.Fatalf("[Segmentio] Failed to dial leader: %s", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		segmentio.Message{
			Key:   []byte(message.Key),
			Value: []byte(message.Value),
		},
	)

	if err != nil {
		log.Fatalf("[Segmentio] Failed to write messages: %s", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatalf("[Segmentio] Failed to close writer: %s", err)
	}

	log.Info("[Segmentio] Message sent")
}

// Confluent library
func Confluent(message Message) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": host})

	if err != nil {
		log.Errorf("[Confluent] Failed to create producer: %s", err)
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *confluent.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				}
			case confluent.Error:
				log.Printf("[Confluent] Error: %v\n", ev)
			default:
				log.Printf("[Confluent] Ignored event: %s\n", ev)
			}
		}
	}()

	err = producer.Produce(&confluent.Message{
		TopicPartition: confluent.TopicPartition{Topic: &topic, Partition: confluent.PartitionAny},
		Value:          []byte(message.Value),
		Key:            []byte(message.Key),
	}, nil)

	if err != nil {
		log.Errorf("[Confluent] Failed to produce message: %s", err)
	}

	producer.Flush(500)

	log.Info("[Confluent] Message sent")
}
