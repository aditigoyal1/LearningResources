package main

import (
	"context"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/verloop/go-tools/kafkamgr"
	"github.com/verloop/go-tools/vcontext"
)

func main() {
	logger := logrus.StandardLogger()
	placeHolderCtx := vcontext.WithLogger(vcontext.WithClientID(nil, "aditi"), logger)

	mgr, err := kafkamgr.New(kafkamgr.DefaultSettings())
	if err != nil {
		logger.WithError(err).Error("failed to init Kafka Manager")
	}

	for i := 0; i < 100; i++ {
		ctx, cancel := context.WithTimeout(placeHolderCtx, 2*time.Second)
		_, err := mgr.Produce(ctx, "msg_produce11", "", []byte("holaaa"))
		cancel()
		if err != nil {
			logger.WithError(err).Error("failed to produce")
		}
		log.Print("Produced a message to msg_topic")
	}
	mgr.Close(context.Background())
}
