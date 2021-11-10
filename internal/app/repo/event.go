package repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/srv-verification-api/internal/model"
)

type EventRepo interface {
	Lock(ctx context.Context, n uint64) ([]model.VerificationEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64) error
	Add(ctx context.Context, event []model.VerificationEvent) error
	Remove(ctx context.Context, eventIDs []uint64) error
}
type repo struct {
	db *sqlx.DB
}

func NewEventRepo(db *sqlx.DB) EventRepo {
	return &repo{db: db}
}

func (r repo) Lock(ctx context.Context, n uint64) ([]model.VerificationEvent, error) {
	query, args, err := sq.Update("verification_events").
		Set("event_status", model.Processed).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"event_status": model.Deferred}).
		Limit(n).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return nil, err
	}

	eventIds := make([]uint64, 0, n)
	err = r.db.SelectContext(ctx, &eventIds, query, args...)

	if err != nil {
		return nil, err
	}

	query, args, err = sq.Select("verification_events.id",
		"verification_events.event_type",
		"verification_events.event_status",
		"verification.id",
		"verification.name").
		Join("verification on verification.id = verification_events.verification_id").
		From("verification_events").
		Where(sq.Eq{"verification_events.id": eventIds}).
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
	return events, nil

}

func (r repo) Unlock(ctx context.Context, eventIDs []uint64) error {
	query, args, err := sq.Update("verification_events").
		PlaceholderFormat(sq.Dollar).
		Set("status", model.Deferred).
		Set("updated_at", "NOW()").
		Where(sq.Eq{"id": eventIDs}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	return err

}

func (r repo) Remove(ctx context.Context, eventIDs []uint64) error {
	query, args, err := sq.Update("verification_events").
		PlaceholderFormat(sq.Dollar).
		Set("status", model.Processed).
		Where(sq.Eq{"id": eventIDs}).
		ToSql()

	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)

	return err
}

func (r repo) Add(ctx context.Context, event []model.VerificationEvent) error {
	return nil
}
