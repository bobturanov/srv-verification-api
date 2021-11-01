package repo

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/srv-verification-api/internal/model"
)

var errNotImplementedMethod = errors.New("method is not implemented")

// Repo is DAO for Verification
type Repo interface {
	DescribeVerification(ctx context.Context, verificationID uint64) (*model.Verification, error)
	AddVerification(ctx context.Context, verification *model.Verification) error
	ListVerification(ctx context.Context) ([]*model.Verification, error)
	RemoveVerification(ctx context.Context, verificationID uint64) (status bool, err error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) *repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeVerification(ctx context.Context, verificationID uint64) (*model.Verification, error) {
	return &model.Verification{ID: 1, Name: "dddd"}, nil
}
