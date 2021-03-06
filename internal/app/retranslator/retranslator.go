package retranslator

import (
	"context"
	"time"

	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"

	"github.com/gammazero/workerpool"
	"github.com/ozonmp/srv-verification-api/internal/app/consumer"
	"github.com/ozonmp/srv-verification-api/internal/app/producer"
	"github.com/ozonmp/srv-verification-api/internal/app/repo"
	"github.com/ozonmp/srv-verification-api/internal/app/sender"
	"github.com/ozonmp/srv-verification-api/internal/model"
)

type Retranslator interface {
	Start(ctx context.Context)
	Close()
}

type Config struct {
	ChannelSize uint64

	ConsumerCount  uint64
	ConsumeSize    uint64
	ConsumeTimeout time.Duration

	ProducerCount uint64
	WorkerCount   int

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan model.VerificationEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
	cancel     context.CancelFunc
}

// NewRetranslator create new retranslator.
func NewRetranslator(cfg Config) Retranslator {
	events := make(chan model.VerificationEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewDbConsumer(
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		events)
	producer := producer.NewKafkaProducer(
		cfg.ProducerCount,
		cfg.Sender,
		cfg.Repo,
		events,
		workerPool)

	return &retranslator{
		events:     events,
		consumer:   consumer,
		producer:   producer,
		workerPool: workerPool,
	}
}

func (r *retranslator) Start(ctx context.Context) {
	logger.InfoKV(ctx, "Retranslator starts")
	ctx, cancel := context.WithCancel(ctx)

	r.cancel = cancel
	r.producer.Start(ctx)
	r.consumer.Start(ctx)
}

func (r *retranslator) Close() {
	logger.InfoKV(context.Background(), "Retranslator stops")
	r.cancel()
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
}
