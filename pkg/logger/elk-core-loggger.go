package logger

import (
	"awesomeProject/pkg/setting"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
)

type ElasticsearchCore struct {
	*zap.Logger
}
type LogstashWriter struct {
	Address string
}

func NewLoggerELk(config setting.LoggerSetting) *ElasticsearchCore{

	var level zapcore.Level
	switch config.Log_level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEnCoderLogELK()


	writer := zapcore.AddSync(&LogstashWriter{
		Address: "localhost:5045",
	})


	core := zapcore.NewCore(encoder, writer, level)


	logger := zap.New(core, zap.AddCaller())

	return &ElasticsearchCore{Logger: logger}
}
func getEnCoderLogELK() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)

}

func (w *LogstashWriter) Write(p []byte) (n int, err error) {
	conn, err := net.Dial("tcp", w.Address)
	if err != nil {
		fmt.Printf("Failed to connect to Logstash at %s: %v\n", w.Address, err)
		return 0, err
	}
	defer conn.Close()

	n, err = conn.Write(p)
	if err != nil {
		fmt.Printf("Failed to send data to Logstash: %v\n", err)
		return n, err
	}
	fmt.Printf("Successfully sent %d bytes to Logstash\n", n)
	return n, nil
}

func (w *LogstashWriter) Sync() error {
	return nil
}
func (l *LoggerZap) SendLogToLogstash(logData interface{}) error {


	data, err := json.Marshal(logData)
	if err != nil {
		return err
	}

	// Kết nối đến Logstash qua TCP
	conn, err := net.Dial("tcp", "localhost:5045")
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}