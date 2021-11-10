package repo

import (
	"context"

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

