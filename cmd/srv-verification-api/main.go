package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ozonmp/srv-verification-api/internal/app/repo"
	"github.com/ozonmp/srv-verification-api/internal/database"

	"github.com/ozonmp/srv-verification-api/internal/app/retranslator"
	"github.com/ozonmp/srv-verification-api/internal/app/sender"
	"github.com/ozonmp/srv-verification-api/internal/config"
	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"
	"github.com/ozonmp/srv-verification-api/internal/tracer"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
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

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfgInst.Database.Host,
		cfgInst.Database.Port,
		cfgInst.Database.User,
		cfgInst.Database.Password,
		cfgInst.Database.Name,
		cfgInst.Database.SslMode)

	db, err := database.NewPostgres(initCtx, dsn, cfgInst.Database.Driver)

	if err != nil {
		logger.FatalKV(ctx, "Failed init postgres", "err", err)
	}

	kafka, err := sender.NewEventProducer(cfgInst.Kafka.Brokers, cfgInst.Kafka.Topic)
	if err != nil {
		logger.FatalKV(ctx, "Failed init sendler", "err", err)
	}

	cfg := retranslator.Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ProducerCount:  28,
		WorkerCount:    2,
		ConsumeTimeout: 10 * time.Second,
		Repo:           repo.NewEventRepo(db),
		Sender:         kafka,
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start(ctx)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	retranslator.Close()
}
