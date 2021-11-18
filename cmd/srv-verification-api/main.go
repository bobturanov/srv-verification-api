package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ozonmp/srv-verification-api/internal/config"
	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"
	"github.com/ozonmp/srv-verification-api/internal/tracer"

	"github.com/ozonmp/srv-verification-api/internal/app/retranslator"
)

func main() {

	sigs := make(chan os.Signal, 1)
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.FatalKV(ctx, "Failed init configuration", "err", err)
	}

	cfgInst := config.GetConfigInstance()

	tracing, err := tracer.NewTracer(&cfgInst)
	if err != nil {
		logger.ErrorKV(ctx, "Failed init tracing", "err", err)
		return
	}
	defer tracing.Close()

	cfg := retranslator.Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ProducerCount:  28,
		WorkerCount:    2,
		ConsumeTimeout: 10 * time.Second, // to run tests
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start(ctx)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	retranslator.Close()
}
