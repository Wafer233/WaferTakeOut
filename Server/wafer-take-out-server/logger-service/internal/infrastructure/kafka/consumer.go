package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MinBytes: 1,    // 最小拉取
		MaxBytes: 10e6, // 最大拉取
	})

	return &Consumer{reader: r}
}

func (c *Consumer) Fetch() (*LogMsg, error) {
	fmt.Println("[Consumer] 等待 Kafka 消息...")
	msg, err := c.reader.ReadMessage(context.Background())
	if err != nil {
		return nil, err
	}

	var logmsg LogMsg
	if err := json.Unmarshal(msg.Value, &logmsg); err != nil {
		return nil, err
	}

	return &logmsg, nil
}
