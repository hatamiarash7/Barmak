package kafka

import (
	"context"
	"time"

	sarama "github.com/Shopify/sarama"
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
	producer, err := confluent.NewProducer(&confluent.ConfigMap{"bootstrap.servers": host})

	if err != nil {
		log.Errorf("[Confluent] Failed to create producer: %s", err)
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *confluent.Message:
				if ev.TopicPartition.Error != nil {
					log.Errorf("[Confluent] Delivery failed: %v\n", ev.TopicPartition)
				}
			case confluent.Error:
				log.Errorf("[Confluent] Error: %v\n", ev)
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

// Sarama library
func Sarama(message Message) {
	producer, err := sarama.NewSyncProducer([]string{host}, nil)

	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Errorf("[Sarama] Failed to create producer: %s", err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(message.Key),
		Value: sarama.StringEncoder(message.Value),
	}

	_, _, err = producer.SendMessage(msg)

	if err != nil {
		log.Errorf("[Sarama] Failed to send message: %s\n", err)
	} else {
		log.Info("[Sarama] Message sent")
	}
}
