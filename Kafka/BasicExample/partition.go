package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func CreateNewTopics(topic string, url string) {
	//topic := "my-topic"

	conn, err := kafka.Dial("tcp", url)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     10,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

func newReader(url string, topic string, partition int) *kafka.Reader {

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{url},
		Topic:   topic,
		//Partition: partition,
		GroupID: "message",
	})
}

func read(url string, topic string, partition int) {

	reader := newReader(url, topic, partition)
	defer reader.Close()
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		log.Printf("rec%d:\t%s\n", msg.Partition, msg.Value)
	}
}

func newWriter(url string, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{url},
		Topic:        topic,
		Balancer:     &kafka.CRC32Balancer{},
		BatchSize:    10,
		BatchTimeout: 1 * time.Millisecond,
	})
}

func write(url string, topic string) {
	writer := newWriter(url, topic)
	defer writer.Close()
	for i := 0; ; i++ {
		v := []byte("V" + strconv.Itoa(i))
		log.Printf("send:\t%s\n", v)
		msg := kafka.Message{Key: v, Value: v}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// func main() {
// 	url := "localhost:9092"
// 	topic := "test-with-partition"
// 	// username := "________________"
// 	// password := "________________"
// 	// clientID := "________________"
// 	// dialer := newDialer(clientID, username, password)
// 	ctx := context.Background()
// 	CreateNewTopics(topic, url)
// 	for i := 0; i < 6; i++ {
// 		go read(url, topic, i)
// 	}

// 	go write(url, topic)
// 	<-ctx.Done()
// }
