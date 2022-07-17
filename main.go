package main

import (
	"arash-hatami.ir/Barmak/kafka"
)

func main() {
	kafka.Segmentio([]kafka.Message{{
		Key:   "Segmentio",
		Value: "hello world",
	}})

	kafka.Confluent([]kafka.Message{{
		Key:   "Confluent",
		Value: "hello world",
	}})

	kafka.Sarama([]kafka.Message{{
		Key:   "Sarama",
		Value: "hello world",
	}})

	kafka.Goka([]kafka.Message{{
		Key:   "Goka",
		Value: "hello world",
	}})
}
