package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "message-log11333"
	brokerAddress = "localhost:9092"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go produce(ctx)
	consume(ctx)
}

func produce(ctx context.Context) {
	// initialize a counter
	i := 0

	l := log.New(os.Stdout, "kafka writer: ", 0)

	// intialize the writer with the broker addresses, and the topic
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Compression:  kafka.Snappy,
		RequiredAcks: kafka.RequireAll,
		Balancer:     kafka.CRC32Balancer{}, // Makes sure msg with a key is sent to the same partition
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		BatchTimeout: 20 * time.Millisecond,
		BatchSize:    10,
		Async:        false,
		Logger:       l,
	}
	time.Sleep(2 * time.Second)
	// w1 := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers: []string{brokerAddress},
	// 	Topic:   topic,
	// 	// assign the logger to the writer
	// 	Logger: l,
	// })

	for {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Topic: topic,

			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}

func consume(ctx context.Context) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	// kafka.NewReader(kafka.ReaderConfig{
	// 	Topic:             s.Topic,
	// 	GroupID:           s.GroupID,
	// 	Brokers:           f.setting.Brokers,
	// 	CommitInterval:    s.CommitInterval,
	// 	StartOffset:       s.Offset,
	// 	ReadBackoffMin:    s.ReadBackoffMin,
	// 	ReadBackoffMax:    s.ReadBackoffMax,
	// 	MinBytes:          s.MinBytes,
	// 	MaxBytes:          s.MaxBytes,
	// 	HeartbeatInterval: 4 * time.Second,
	// })
	r := kafka.NewReader(kafka.ReaderConfig{
		StartOffset:       kafka.FirstOffset,
		Brokers:           []string{brokerAddress},
		Topic:             topic,
		GroupID:           "my-group1",
		HeartbeatInterval: 4 * time.Second,
		MinBytes:          10e3, // 10KB
		MaxBytes:          10e6, // 10MB
		// assign the logger to the reader
		Logger: l,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}
