package consumer

import (
	"context"
	"srv-verification-api/internal/app/repo"
	"srv-verification-api/internal/model"
	"sync"
	"time"

)

type Consumer interface {
	Start()
	Close()
}

type consumer struct {
	n      uint64
	events chan<- model.VerificationEvent
	repo repo.EventRepo
	batchSize uint64
	timeout   time.Duration
	cxt context.Context
	wg   *sync.WaitGroup
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
	events chan<- model.VerificationEvent, cxt context.Context) Consumer {

	wg := &sync.WaitGroup{}

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		events:    events,
		wg:        wg,
		cxt:      cxt,
	}
}

func (c *consumer) Start() {
	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)
			go c.processConsumer()
	}
}

func (c *consumer) Close() {
	c.wg.Wait()
}

func (c *consumer) processConsumer() {
	defer c.wg.Done()
	ticker := time.NewTicker(c.timeout)
	for {
		select {
		case <-ticker.C:
			c.processEvent()
		case <-c.cxt.Done():
			return
		}
	}
}

func (c *consumer) processEvent() {
	events, err := c.repo.Lock(c.batchSize)
	if err != nil {
		return
	}
	unlock := make([]uint64, c.batchSize)
	for i, event := range events {
		event.Type = model.Updated
		unlock[i] = event.ID
		c.events <- event
	}
}
