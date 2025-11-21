package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// cli的美化，加上颜色， 本质上就是自定义encoderLevel

const (
	Blue   = "\033[34m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
)

// 自定义 encodeLevel
// info -> blue 等等
func encodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		enc.AppendString(Blue + "INFO" + Reset)
	case zapcore.WarnLevel:
		enc.AppendString(Yellow + "WARN" + Reset)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(Red + "ERROR" + Reset)
	default:
		enc.AppendString(level.String())
	}
}

// 给整条日志加前缀
type MyEncoder struct {
	zapcore.Encoder
}

func (m MyEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := m.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return buf, err
	}

	str := buf.String()
	buf.Reset()
	buf.AppendString("[api-gateway] ")
	buf.AppendString(str)

	return buf, nil
}

func NewConsoleEncoder() zapcore.Encoder {
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeLevel = encodeLevel
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	return MyEncoder{
		zapcore.NewConsoleEncoder(cfg),
	}
}

func NewFileEncoder() zapcore.Encoder {
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	return MyEncoder{
		zapcore.NewConsoleEncoder(cfg),
	}
}
