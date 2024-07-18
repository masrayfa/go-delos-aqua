package helper

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CommitOrRollback(ctx context.Context, tx pgx.Tx) error {
	err := recover()
	if err != nil {
		err := tx.Rollback(ctx)
		return err
	} else {
		err := tx.Commit(ctx)
		return err
	}
}