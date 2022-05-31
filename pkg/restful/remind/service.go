package remind

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
)

type service struct {
	repo iface.IRepository
}

func (s *service) ListRemind(ctx context.Context, params model.ListRemindParams, filterParams types.FilterRemindParams) (reminds []model.Remind, count int64, err error) {
	reminds = make([]model.Remind, 0)
	err = s.repo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		var err2 error

		_, err2 = tx.Exec("SET @content = ?", filterParams.Content)
		if err2 != nil {
			return err2
		}

		_, err2 = tx.Exec("SET @status = ?", filterParams.Status)
		if err2 != nil {
			return err2
		}

		reminds, err2 = s.repo.WithTx(tx).ListRemind(ctx, params)
		if err2 != nil {
			return err2
		}

		count, err2 = s.repo.WithTx(tx).CountListRemind(ctx)
		if err2 != nil {
			return err2
		}

		return nil
	})

	return
}

func (s *service) GetRemind(ctx context.Context, remindId int64) (model.Remind, error) {
	return s.repo.GetRemind(ctx, remindId)
}

func (s *service) CreateRemind(ctx context.Context, params model.CreateRemindParams) error {
	return s.repo.CreateRemind(ctx, params)
}

func (s *service) UpdateRemind(ctx context.Context, params model.UpdateRemindParams) error {
	return s.repo.UpdateRemind(ctx, params)
}

func (s *service) DeleteRemind(ctx context.Context, remindId int64) error {
	return s.repo.DeleteRemind(ctx, remindId)
}

func (s *service) ListActiveRemind(ctx context.Context) ([]model.ListActiveRemindRow, error) {
	return s.repo.ListActiveRemind(ctx)
}

func NewService(Repo iface.IRepository) iface.IRemindService {
	return &service{
		repo: Repo,
	}
}
