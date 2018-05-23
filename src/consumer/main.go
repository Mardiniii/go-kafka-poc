package consumer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Start triggers a new consumer routine
func Start() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "pocGroup",
		"auto.offset.reset": "earliest",
	})
	failOnError(err, "Consumer cannot be created")

	file := createFile()
	defer file.Close()
	c.SubscribeTopics([]string{"pocTopic", "^aRegex.*[Tt]opic"}, nil)

	for {
		// -1 for indefinite wait or timeout
		msg, err := c.ReadMessage(-1)
		if err == nil {
			file.WriteString(string(msg.Value) + "\n")
			fmt.Printf("Saved on file %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}
