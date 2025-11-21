package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

//把日志封装成 JSON
//
//通过 Producer 发到 Kafka
//
//做的就是 "send to Kafka" 的事

// Producer 封装 Kafka Writer（生产者）
// 可以在任何地方注入（service / logger 等）
type Producer struct {
	writer *kafka.Writer
}

// NewProducer 创建一个 Kafka Producer
func NewProducer(addr string) *Producer {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(addr),
		Balancer:               &kafka.LeastBytes{}, // 更均衡的 partition 分配
		RequiredAcks:           kafka.RequireAll,    // 确保 Leader + Follower 都写入
		Async:                  false,               // 关闭异步（交给我们自己的 channel 控制）
		AllowAutoTopicCreation: true,                // 需要时自动创建 topic
	}

	return &Producer{
		writer: w,
	}
}

// Send 给指定 Topic 发送消息（同步发送）
// 注意：上层必须包装成异步（例如 channel 模式）
// 这样可以保持 Producer 简单而可复用
func (p *Producer) Send(topic string, key []byte, value []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	msg := kafka.Message{
		Key:   key,
		Topic: topic,
		Value: value,
	}

	return p.writer.WriteMessages(ctx, msg)
}

// Close 关闭 producer（程序退出时调用）
func (p *Producer) Close() error {
	return p.writer.Close()
}
