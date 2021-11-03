package producer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gammazero/workerpool"
	"github.com/ozonmp/srv-verification-api/internal/app/repo"
	"github.com/ozonmp/srv-verification-api/internal/app/sender"
	"github.com/ozonmp/srv-verification-api/internal/model"
)

type Producer interface {
	Start(ctx context.Context)
	Close()
}

type producer struct {
	producerCount uint64
	timeout       time.Duration

	sender sender.EventSender
	events <-chan model.VerificationEvent

	workerPool *workerpool.WorkerPool

	wg   *sync.WaitGroup
	repo repo.EventRepo
}

func NewKafkaProducer(
	producerCount uint64,
	sender sender.EventSender,
	repo repo.EventRepo,
	events <-chan model.VerificationEvent,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}

	return &producer{
		producerCount: producerCount,
		sender:        sender,
		repo:          repo,
		events:        events,
		workerPool:    workerPool,
		wg:            wg,
	}
}

func (p *producer) Start(ctx context.Context) {
	for i := uint64(0); i < p.producerCount; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					p.produceEvent(&event)
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	p.wg.Wait()
}

func (p *producer) produceEvent(event *model.VerificationEvent) {
	if err := p.sender.Send(event); err != nil {
		log.Printf("Error sending Event ID: %d to Kafka", event.ID)
		p.workerPool.Submit(func() {
			if err := p.repo.Unlock([]uint64{event.ID}); err != nil {
				log.Printf("Error unlocking Event ID: %d in DB", event.ID)
			}
		})
	} else {
		p.workerPool.Submit(func() {
			if err := p.repo.Remove([]uint64{event.ID}); err != nil {
				log.Printf("Error removing Event ID: %d", event.ID)
			}
		})
	}
}
