package repository

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg/interface"
	"database/sql"
)

func (r *repository) WithTx(tx *sql.Tx) model.Querier {
	return r.Queries.WithTx(tx)
}

func (r *repository) Transaction(ctx context.Context, cb iface.Callback) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if err = cb(ctx, tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()
	return nil
}
