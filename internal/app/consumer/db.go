package consumer

import (
	"context"
	"sync"
	"time"

	"github.com/ozonmp/srv-verification-api/internal/app/repo"
	"github.com/ozonmp/srv-verification-api/internal/model"
)

type Consumer interface {
	Start(ctx context.Context)
	Close()
}

type consumer struct {
	n         uint64
	events    chan<- model.VerificationEvent
	repo      repo.EventRepo
	batchSize uint64
	timeout   time.Duration
	wg        *sync.WaitGroup
}

type Config struct {
	n         uint64
	events    chan<- model.VerificationEvent
	repo      repo.EventRepo
	batchSize uint64
	timeout   time.Duration
}

func NewDbConsumer(
	n uint64,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.EventRepo,
	events chan<- model.VerificationEvent) Consumer {

	wg := &sync.WaitGroup{}

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		events:    events,
		wg:        wg,
	}
}

func (c *consumer) Start(ctx context.Context) {
	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)
		go c.processConsumer(ctx)
	}
}

func (c *consumer) Close() {
	c.wg.Wait()
}

func (c *consumer) processConsumer(ctx context.Context) {
	defer c.wg.Done()
	ticker := time.NewTicker(c.timeout)
	for {
		select {
		case <-ticker.C:
			c.processEvent(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (c *consumer) processEvent(ctx context.Context) {
	events, err := c.repo.Lock(ctx, c.batchSize)
	if err != nil {
		return
	}
	unlock := make([]uint64, c.batchSize)
	for i, event := range events {
		event.Type = model.Updated
		unlock[i] = event.ID
		c.events <- event
		repo.TotalHandledEvents.Add(1)
	}
}
