package main

import (
	"context"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"

	"github.com/verloop/go-tools/kafkamgr"
)

func main() {
	// Inits Kafka Manager with default setting
	// It connects to localhost with default setting
	mgr, err := kafkamgr.New(kafkamgr.DefaultSettings())
	if err != nil {
		panic(err)
	}

	consumer := kafkamgr.DefaultConsumerSettings("msg_topic", handleMsg)
	consumer.Offset = kafka.LastOffset
	if err := mgr.AddConsumer(consumer); err != nil {
		logrus.Fatal(err)
	}

	// Handle graceful shutdown
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-stop
		mgr.Close(context.Background())
	}()

	// blocking call
	mgr.Consume()
}

// Callback handler for consumer to call
// Handler is responsible for reading the msg, this is  a blocking call
func handleMsg(reader *kafka.Reader, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
	loopy:
		msg, err := reader.FetchMessage(context.Background())
		if err != nil {
			if err == io.EOF {
				return
			}
			logrus.WithError(err).Error("failed to fetch message")
			goto loopy
		}

		logrus.WithField("msg", string(msg.Value)).Infof("received message from topic[%s] with partition[%d] & offset[%d]", msg.Topic, msg.Partition, msg.Offset)
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		if err := reader.CommitMessages(ctx, msg); err != nil {
			logrus.WithError(err).Error("failed to commit msgs")
			cancel()
			goto loopy
		}
		cancel()
	}
}
