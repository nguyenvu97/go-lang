package kafka

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	kafka "github.com/segmentio/kafka-go"
//	"strings"
//	"time"
//)
//
//var (
//	KafkaProducer *kafka.Writer
//)
//
//const (
//	kafkaURL = "localhost:9092"
//	topic    = "go_topic"
//)
//
//func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
//	return &kafka.Writer{
//		Addr:     kafka.TCP(kafkaURL),
//		Topic:    topic,
//		Balancer: &kafka.LeastBytes{},
//	}
//}
//
//func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
//	brokers := strings.Split(kafkaURL, ",")
//	return kafka.NewReader(kafka.ReaderConfig{
//		Brokers:        brokers,
//		GroupID:        groupID,
//		Topic:          topic,
//		MinBytes:       10e3, // 10KB
//		MaxBytes:       10e6, // 10MB
//		CommitInterval: time.Second,
//		StartOffset:    kafka.FirstOffset,
//	})
//}
//
//type StockInfo struct {
//	Msg  string `json:"msg"`
//	Type string `json:"type"`
//}
//
//func newStock(msg, typeMsg string) *StockInfo {
//	s := StockInfo{msg, typeMsg}
//	return &s
//}
//
//func actionStock(g *gin.Context) {
//	s := newStock(g.Query("msg"), g.Query("type"))
//	body := make(map[string]interface{})
//	body["info"] = s
//
//	jsonBody, _ := json.Marshal(body)
//	msg := kafka.Message{
//		Key:   []byte("action"),
//		Value: []byte(jsonBody),
//	}
//	err := KafkaProducer.WriteMessages(context.Background(), msg)
//	if err != nil {
//		g.JSON(200, gin.H{
//			"error": err.Error(),
//		})
//		return
//	}
//
//	g.JSON(200, gin.H{
//		"error": "",
//		"msg":   "send meg ok",
//	})
//	return
//
//}
//func registerConsumer(id int) {
//	kafkaGroupId := "consumer-group-"
//	render := getKafkaReader(kafkaURL, topic, kafkaGroupId)
//	defer render.Close()
//	fmt.Printf("cumsumer (%d) hong phien atc ", id)
//
//	for {
//		m, err := render.ReadMessage(context.Background())
//		if err != nil {
//			fmt.Printf("cumsumer (%d) err(%err) ", id, err)
//		}
//
//		fmt.Printf("cunsumer (%m) , tipic(%t),pattion (%p), offser(%s), time := %d %s : %d\n", id, topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
//	}
//
//}
//
////func main() {
////	r := gin.Default()
////	KafkaProducer = getKafkaWriter(kafkaURL, topic)
////	defer KafkaProducer.Close()
////
////	r.POST("action/stock", actionStock)
////
////	go registerConsumer(1)
////	go registerConsumer(2)
////
////	//r.Run(":1997")
////}
