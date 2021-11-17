package repo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/srv-verification-api/internal/model"
	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"
)

func (r repo) Remove(ctx context.Context, eventIDs []uint64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.Remove")
	defer span.Finish()

	query, args, err := squirrel.Update("verification_events").
		PlaceholderFormat(squirrel.Dollar).
		Set("status", model.Processed).
		Where(squirrel.Eq{"id": eventIDs}).
		ToSql()

	if err != nil {
		logger.ErrorKV(ctx, "repo.Remove() get result query", "err", err)
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)

	return err
}
