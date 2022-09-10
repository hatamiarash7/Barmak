package kafka

import (
	"context"
	"time"

	sarama "github.com/Shopify/sarama"
	confluent "github.com/confluentinc/confluent-kafka-go/kafka"
	goka "github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	segmentio "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var (
	ctx       context.Context = context.Background()
	host      string          = "localhost:9092"
	partition int             = 0
	topic     string          = "benchmark"
)

// Message global message structure for kafka
type Message struct {
	Key   string
	Value string
}

// Segmentio library
func Segmentio(messages []Message) {
	conn, err := segmentio.DialLeader(context.Background(), "tcp", host, topic, partition)

	if err != nil {
		log.Fatalf("[Segmentio] Failed to dial leader: %s", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	for _, message := range messages {
		_, err = conn.WriteMessages(
			segmentio.Message{
				Key:   []byte(message.Key),
				Value: []byte(message.Value),
			},
		)

		if err != nil {
			log.Fatalf("[Segmentio] Failed to write messages: %s", err)
		}
	}

	if err := conn.Close(); err != nil {
		log.Fatalf("[Segmentio] Failed to close writer: %s", err)
	}
}

// Confluent library
func Confluent(messages []Message) {
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

	for _, message := range messages {
		err = producer.Produce(&confluent.Message{
			TopicPartition: confluent.TopicPartition{Topic: &topic, Partition: confluent.PartitionAny},
			Value:          []byte(message.Value),
			Key:            []byte(message.Key),
		}, nil)

		if err != nil {
			log.Errorf("[Confluent] Failed to produce message: %s", err)
		}

		producer.Flush(1)
	}
}

// Sarama library
func Sarama(messages []Message) {
	producer, err := sarama.NewAsyncProducer([]string{host}, nil)

	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Errorf("[Sarama] Failed to create producer: %s", err)
		}
	}()

	for _, message := range messages {
		producer.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(message.Key),
			Value: sarama.StringEncoder(message.Value),
		}

		if err != nil {
			log.Errorf("[Sarama] Failed to send message: %s\n", err)
		}
	}
}

// Goka library
func Goka(messages []Message) {
	emitter, err := goka.NewEmitter(
		[]string{host},
		goka.Stream(topic),
		new(codec.String),
	)

	if err != nil {
		log.Fatalf("[Goka] Error creating emitter: %v", err)
	}

	defer emitter.Finish()

	for _, message := range messages {
		err = emitter.EmitSync(message.Key, message.Value)

		if err != nil {
			log.Fatalf("[Goka] Error emitting message: %v", err)
		}
	}
}
