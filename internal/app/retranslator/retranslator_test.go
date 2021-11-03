package retranslator

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/srv-verification-api/internal/mocks"
	"github.com/ozonmp/srv-verification-api/internal/model"
	"testing"
	"time"
)

var eventData = []model.VerificationEvent{
	{ID: 1, Type: model.Created, Status: model.Processed, Entity: &model.Verification{ID: 1}},
	{ID: 2, Type: model.Created, Status: model.Processed, Entity: &model.Verification{ID: 2}},
	{ID: 3, Type: model.Created, Status: model.Processed, Entity: &model.Verification{ID: 3}},
}

func TestCaseSendAndUnlockProducer(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	gomock.InOrder(
		repo.EXPECT().Lock(uint64(5)).Return(eventData, nil).MinTimes(1).MaxTimes(1),
		sender.EXPECT().Send(&eventData[0]).Return(nil).MinTimes(1).MaxTimes(1),
		repo.EXPECT().Unlock([]uint64{eventData[0].ID}).Return(nil).MinTimes(1).MaxTimes(1),
	)

	startTest(repo, sender)
}

func TestCaseSendErrorAndRemoveProducer(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	gomock.InOrder(
		repo.EXPECT().Lock(uint64(10)).Return(eventData, nil).MinTimes(1).MaxTimes(1),
		sender.EXPECT().Send(&eventData[0]).Return(errors.New("sending is NOT allowed")).MinTimes(1).MaxTimes(1),
		repo.EXPECT().Remove([]uint64{eventData[0].ID}).Return(nil).MinTimes(1).MaxTimes(1),
	)

	startTest(repo, sender)
}

func startTest(repo *mocks.MockEventRepo, sender *mocks.MockEventSender) {

	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: 10 * time.Second,
		ProducerCount:  2,
		WorkerCount:    2,
		Repo:           repo,
		Sender:         sender,
	}
	ctx := context.Background()

	retranslator := NewRetranslator(cfg)
	retranslator.Start(ctx)
	time.Sleep(time.Second / 4)
	retranslator.Close()
}
