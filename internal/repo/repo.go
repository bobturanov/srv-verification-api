package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/opentracing/opentracing-go"

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
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("verificationID", verificationID)
	query, args, err := sq.Select("*").PlaceholderFormat(sq.Dollar).From("verification").Where(sq.Eq{"id": verificationID}).ToSql()

	if err != nil {
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
	queryIDSeq, argsIDSeq, errIDSeq := sq.Select("nextval('verification_id_seq')").ToSql()

	if errIDSeq != nil {
		return errIDSeq
	}

	err := r.db.QueryRowContext(ctx, queryIDSeq, argsIDSeq...).Scan(&verification.ID)

	span := opentracing.SpanFromContext(ctx)
	span.SetTag("verificationID", verification.ID)

	if err != nil {
		return err
	}
	query, args, err := sq.Insert("verification").
		PlaceholderFormat(sq.Dollar).
		Columns("id", "name", "created_at", "updated_at").
		Values(verification.ID, verification.Name, verification.CreatedAt, verification.UpdatedAt).
		Suffix("RETURNING id").ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}
	return nil
}

func (r *repo) ListVerification(ctx context.Context) ([]*model.Verification, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.ListVerification")
	defer span.Finish()
	query, args, err := sq.Select("*").PlaceholderFormat(sq.Dollar).From("verification").ToSql()

	if err != nil {
		return nil, err
	}

	verification := make([]*model.Verification, 0)
	err = r.db.SelectContext(ctx, &verification, query, args...)

	switch err {
	case nil:
		return verification, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, errInternalMethod
	}
}

func (r *repo) RemoveVerification(ctx context.Context, verificationID uint64) (status bool, err error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("verificationID", verificationID)
	query, args, err := sq.Delete("verification").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": verificationID}).ToSql()

	if err != nil {
		return false, err
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	switch err {
	case nil:
		return true, nil
	default:
		return false, errInternalMethod
	}
}
