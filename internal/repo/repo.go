package repo

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/srv-verification-api/internal/model"
)

var ErrNotFound = errors.New("no rows in database for the given query")
var errInternalMethod = errors.New("internal error")

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
func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeVerification(ctx context.Context, verificationID uint64) (*model.Verification, error) {
	query, args, err := sq.Select("*").PlaceholderFormat(sq.Dollar).From("verification").Where(sq.Eq{"id": verificationID}).ToSql()

	if err != nil{
		return nil, err
	}
	verification := model.Verification{}
	err = r.db.GetContext(ctx, &verification, query, args...)

	switch err {
	case nil:
		return &verification, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, errInternalMethod
	}

}

func (r *repo) AddVerification(ctx context.Context, verification *model.Verification) error {
	return errInternalMethod
}

func (r *repo) ListVerification(ctx context.Context) ([]*model.Verification, error) {
	return nil, errInternalMethod
}

func (r *repo) RemoveVerification(ctx context.Context, verificationID uint64) (status bool, err error) {
	return false, errInternalMethod
}
