package logger

import (
	"encoding/json"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/infrastructure/kafka"
	"go.uber.org/zap"
)

var globalKafkaLog *KafkaLog

func InitKafkaLog(p *kafka.Producer) {
	globalKafkaLog = NewKafkaLog(p)
}

// 全局logger
func K() *KafkaLog {
	return globalKafkaLog
}

type KafkaLog struct {
	producer *kafka.Producer
	ch       chan logMsg
}

type logMsg struct {
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	Time    string                 `json:"time"`
	Fields  map[string]interface{} `json:"fields"`
}

func NewKafkaLog(p *kafka.Producer) *KafkaLog {
	kl := &KafkaLog{
		producer: p,
		ch:       make(chan logMsg, 1000),
	}
	go kl.loop()
	return kl
}

func (kl *KafkaLog) loop() {
	for msg := range kl.ch {
		b, _ := json.Marshal(msg)
		err := kl.producer.Send("wafer-take-out-log", []byte("api-gateway日志"), b)
		if err != nil {
			zap.L().Error("kafka日志错误", zap.Error(err))
		}
	}
}

// 结构化数据
func (kl *KafkaLog) Info(msg string, fields ...zap.Field) {
	kl.push("INFO", msg, fields...)
}

func (kl *KafkaLog) Error(msg string, fields ...zap.Field) {
	kl.push("ERROR", msg, fields...)
}

func (kl *KafkaLog) push(level string, message string, fields ...zap.Field) {

	fieldMap := make(map[string]interface{})
	for _, f := range fields {
		fieldMap[f.Key] = f.Interface
	}

	select {
	case kl.ch <- logMsg{
		Level:   level,
		Message: message,
		Fields:  fieldMap,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}:
	default:
		// channel 满 → 丢弃（保护业务）
	}
}
