package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ozonmp/srv-verification-api/internal/app/reader"
	"github.com/ozonmp/srv-verification-api/internal/config"
	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"
	"github.com/ozonmp/srv-verification-api/internal/tracer"
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

	initCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	kafka, err := reader.NewEventConsumer(cfgInst.Kafka.Brokers, cfgInst.Kafka.Topic, cfgInst.Kafka.GroupID)
	if err != nil {
		logger.FatalKV(ctx, "Failed init reader", "err", err)
	}

	kafka.Read(initCtx)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

}
