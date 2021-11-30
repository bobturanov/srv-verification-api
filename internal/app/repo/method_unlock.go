package repo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/srv-verification-api/internal/model"
	"github.com/ozonmp/srv-verification-api/internal/pkg/logger"
)

func (r repo) Unlock(ctx context.Context, eventIDs []uint64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.Unlock")
	defer span.Finish()

	query, args, err := squirrel.Update("verification_events").
		PlaceholderFormat(squirrel.Dollar).
		Set("status", model.Deferred).
		Set("updated_at", "NOW()").
		Where(squirrel.Eq{"event_id": eventIDs}).
		ToSql()

	if err != nil {
		logger.ErrorKV(ctx, "repo.Unlock() get result query", "err", err)
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)

	return err

}
