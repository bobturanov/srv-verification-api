package reader

import (
	"context"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

type EventReader interface {
	Read() error
}

type kafkaReader struct {
	consumer sarama.ConsumerGroup
	topic    []string
	groupId  string
}

type Consumer struct {
	ready chan bool
}

func NewEventConsumer(brokers []string, topic string, groupId string) (*kafkaReader, error) {
	config := sarama.NewConfig()

	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumerGroup(brokers, groupId, config)

	return &kafkaReader{
		consumer: consumer,
		topic:    []string{topic},
		groupId:  groupId,
	}, err
}

func (kr *kafkaReader) Read(ctx context.Context) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	consumer := Consumer{
		ready: make(chan bool),
	}

	go func() {
		defer wg.Done()
		for {

			if err := kr.consumer.Consume(ctx, kr.topic, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}

			consumer.ready = make(chan bool)
		}
	}()

}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for message := range claim.Messages() {
		log.Printf(string(message.Value))
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}

	return nil
}
