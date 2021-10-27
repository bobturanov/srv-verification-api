package main

import (
	"context"
	"os"
	"os/signal"
	"srv-verification-api/internal/app/retranslator"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)
	ctx := context.Background()

	cfg := retranslator.Config{
		ChannelSize:   512,
		ConsumerCount: 2,
		ConsumeSize:   10,
		ProducerCount: 28,
		WorkerCount:   2,
		ConsumeTimeout: 10 * time.Second, // to run tests
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start(ctx)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}