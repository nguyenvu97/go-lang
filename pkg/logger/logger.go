package logger

import (
	"awesomeProject/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type LoggerZap struct {
	*zap.Logger
}
//var (
//	host = "http://localhost:5044"
//)
func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Log_level

	var level zapcore.Level
	switch logLevel {
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

	encoder := getEnCoderLog()

	hook := lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size,
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age,
		Compress:   config.Compress,
	}
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), level)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}
func getEnCoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)

}

//func (l *LoggerZap) SendLogToLogstash(dto interface{}) error {
//	logMessage := map[string]interface{}{
//		"action":  "CreateProduct",
//		"status":  "error",
//		"product": dto,
//	}
//
//	// Convert log message to JSON
//	data, err := json.Marshal(logMessage)
//	if err != nil {
//		l.Error("Failed to marshal log message", zap.Error(err))
//		return err
//	}
//
//	// Kết nối đến Logstash qua TCP
//	conn, err := net.Dial("tcp", "localhost:5044")
//	fmt.Printf("connect %d",conn)
//	if err != nil {
//		l.Error("Failed to connect to Logstash", zap.Error(err))
//		return err
//	}
//	defer conn.Close()
//
//	// Gửi dữ liệu qua TCP
//	_, err = conn.Write(data)
//	if err != nil {
//		l.Error("Failed to send log to Logstash", zap.Error(err))
//		return err
//	}
//
//	return nil
//}
