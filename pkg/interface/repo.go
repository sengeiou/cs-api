package iface

import (
	"context"
	"cs-api/db/model"
	"database/sql"
)

type Callback func(ctx context.Context, tx *sql.Tx) error

type IRepository interface {
	model.Querier
	WithTx(tx *sql.Tx) model.Querier
	Transaction(ctx context.Context, f Callback) error
}
