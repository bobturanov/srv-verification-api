package retranslator

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/srv-verification-api/internal/mocks"
	"github.com/ozonmp/srv-verification-api/internal/model"
)

var eventData = []model.VerificationEvent{
	{ID: 1, Type: model.Created, Status: model.Processed, Entity: &model.Verification{ID: 1}},
	{ID: 2, Type: model.Created, Status: model.Processed, Entity: &model.Verification{ID: 2}},
	{ID: 3, Type: model.Created, Status: model.Processed, Entity: &model.Verification{ID: 3}},
}

func TestCaseSendAndUnlockProducer(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	gomock.InOrder(
		repo.EXPECT().Lock(ctx, uint64(5)).Return(eventData, nil).MinTimes(1).MaxTimes(1),
		sender.EXPECT().Send(&eventData[0]).Return(nil).MinTimes(1).MaxTimes(1),
		repo.EXPECT().Unlock(ctx, []uint64{eventData[0].ID}).Return(nil).MinTimes(1).MaxTimes(1),
	)

	startTest(ctx, repo, sender)
}

func TestCaseSendErrorAndRemoveProducer(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	gomock.InOrder(
		repo.EXPECT().Lock(ctx, uint64(10)).Return(eventData, nil).MinTimes(1).MaxTimes(1),
		sender.EXPECT().Send(&eventData[0]).Return(errors.New("sending is NOT allowed")).MinTimes(1).MaxTimes(1),
		repo.EXPECT().Remove(ctx, []uint64{eventData[0].ID}).Return(nil).MinTimes(1).MaxTimes(1),
	)

	startTest(ctx, repo, sender)
}

func startTest(ctx context.Context, repo *mocks.MockEventRepo, sender *mocks.MockEventSender) {

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

	retranslator := NewRetranslator(cfg)
	retranslator.Start(ctx)
	time.Sleep(time.Second / 4)
	retranslator.Close()
}
