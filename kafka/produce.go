package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
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
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	if err != nil {
		log.Fatal("[Segmentio] Failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   []byte(message.Key),
			Value: []byte(message.Value),
		},
	)

	if err != nil {
		log.Fatal("[Segmentio] Failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("[Segmentio] Failed to close writer:", err)
	}

	log.Info("[Segmentio] Message sent")
}
