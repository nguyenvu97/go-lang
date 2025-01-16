package initialize

import (
	"awesomeProject/global"
	"github.com/segmentio/kafka-go"
	"strings"
	"time"
)

func InitKafka() {
	getKafkaWriter()
	getKafkaReader()

}
func getKafkaWriter() *kafka.Writer {
	k := global.Config.Kafka
	return &kafka.Writer{
		Addr:     kafka.TCP(k.Url),
		Topic:    k.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func getKafkaReader() *kafka.Reader {
	k := global.Config.Kafka
	brokers := strings.Split(k.Url, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        k.Consumer,
		Topic:          k.Topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}
