package repo

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/ozonmp/srv-verification-api/internal/model"
)

func (r repo) Remove(ctx context.Context, eventIDs []uint64) error {
	query, args, err := squirrel.Update("verification_events").
		PlaceholderFormat(squirrel.Dollar).
		Set("status", model.Processed).
		Where(squirrel.Eq{"id": eventIDs}).
		ToSql()

	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, query, args...)

	return err
}

