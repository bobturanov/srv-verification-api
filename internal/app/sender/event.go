package sender

import (
	"github.com/Shopify/sarama"
	"github.com/ozonmp/srv-verification-api/internal/model"
	pb "github.com/ozonmp/srv-verification-api/pkg/srv-verification-api"
	"google.golang.org/protobuf/proto"
)

const producerPartitionerStrategy = -1

type EventSender interface {
	Send(verification *model.VerificationEvent) error
}

type kafkaSender struct {
	producer sarama.SyncProducer
	topic    string
}

func NewEventProducer(brokers []string, topic string) (*kafkaSender, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)

	return &kafkaSender{
		producer: producer,
		topic:    topic,
	}, err
}

func (ks *kafkaSender) Send(event *model.VerificationEvent) error {
	pbVerificationEntity := &pb.VerificationEntity{}

	pbVerificationEntity.VerificationId = event.Entity.ID
	pbVerificationEntity.Name = event.Entity.Name

	pbEvent := &pb.VerificationEvent{
		Id:             event.ID,
		VerificationId: event.VerificationID,
		EventStatus:    string(event.Status),
		EventType:      string(event.Type),
		Entity:         pbVerificationEntity,
	}

	message, err := proto.Marshal(pbEvent)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     ks.topic,
		Partition: producerPartitionerStrategy,
		Value:     sarama.ByteEncoder(message),
	}
	_, _, err = ks.producer.SendMessage(msg)

	return err
}
