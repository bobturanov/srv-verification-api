package repo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/srv-verification-api/internal/model"
	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"
)

func (r repo) Lock(ctx context.Context, n uint64) ([]model.VerificationEvent, error) {
	spanLock, ctx := opentracing.StartSpanFromContext(ctx, "repo.Lock")
	defer spanLock.Finish()

	eventIds, err := r.getEventIdsFromDB(ctx, n)

	if err != nil {
		logger.ErrorKV(ctx, "repo.getEventIdsFromDB() get result query", "err", err)
		return nil, err
	}

	eventsData, err := r.getEventsDataFromDB(ctx, err, eventIds)

	if err != nil {
		logger.ErrorKV(ctx, "repo.getEventsDataFromDB() get result query", "err", err)
		return nil, err
	}

	events := r.convertToVerificationEventModel(n, eventsData)
	return events, nil

}

func (r repo) convertToVerificationEventModel(n uint64, eventsData []struct {
	EventId          uint64
	EventType        model.EventType
	EventStatus      model.EventStatus
	VerificationId   uint64
	VerificationName string
}) []model.VerificationEvent {

	events := make([]model.VerificationEvent, 0, n)

	for _, event := range eventsData {
		events = append(events, model.VerificationEvent{
			ID:     event.EventId,
			Type:   event.EventType,
			Status: event.EventStatus,
			Entity: &model.Verification{
				ID:   event.VerificationId,
				Name: event.VerificationName,
			},
		})
	}
	return events
}

func (r repo) getEventsDataFromDB(ctx context.Context, err error, eventIds []uint64) ([]struct {
	EventId          uint64
	EventType        model.EventType
	EventStatus      model.EventStatus
	VerificationId   uint64
	VerificationName string
}, error) {

	query, args, err := squirrel.Select("verification_events.id",
		"verification_events.event_type",
		"verification_events.event_status",
		"verification.id",
		"verification.name").
		Join("verification on verification.id = verification_events.verification_id").
		From("verification_events").
		Where(squirrel.Eq{"verification_events.id": eventIds}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var eventsData []struct {
		EventId          uint64
		EventType        model.EventType
		EventStatus      model.EventStatus
		VerificationId   uint64
		VerificationName string
	}

	err = r.db.SelectContext(ctx, &eventsData, query, args...)

	if err != nil {
		return nil, err
	}

	return eventsData, nil
}

func (r repo) getEventIdsFromDB(ctx context.Context, n uint64) ([]uint64, error) {
	query, args, err := squirrel.Update("verification_events").
		Set("event_status", model.Processed).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"event_status": model.Deferred}).
		Limit(n).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		logger.ErrorKV(ctx, "repo.getEventIdsFromDB() get select query", "err", err)
		return nil, err
	}

	eventIds := make([]uint64, 0, n)
	err = r.db.SelectContext(ctx, &eventIds, query, args...)

	if err != nil {
		logger.ErrorKV(ctx, "repo.getEventIdsFromDB() get result query", "err", err)
		return nil, err
	}
	return eventIds, nil
}
