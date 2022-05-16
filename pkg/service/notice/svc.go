package notice

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg/types"
	"database/sql"
)

func (s *service) ListNotice(ctx context.Context, params model.ListNoticeParams, filterParams types.FilterNoticeParams) (notices []model.Notice, count int64, err error) {
	notices = make([]model.Notice, 0)
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

		notices, err2 = s.repo.WithTx(tx).ListNotice(ctx, params)
		if err2 != nil {
			return err2
		}

		count, err2 = s.repo.WithTx(tx).CountListNotice(ctx)
		if err2 != nil {
			return err2
		}

		return nil
	})

	return
}

func (s *service) GetNotice(ctx context.Context, noticeId int64) (model.Notice, error) {
	return s.repo.GetNotice(ctx, noticeId)
}

func (s *service) CreateNotice(ctx context.Context, params model.CreateNoticeParams) error {
	return s.repo.CreateNotice(ctx, params)
}

func (s *service) UpdateNotice(ctx context.Context, params model.UpdateNoticeParams) error {
	return s.repo.UpdateNotice(ctx, params)
}

func (s *service) DeleteNotice(ctx context.Context, noticeId int64) error {
	return s.repo.DeleteNotice(ctx, noticeId)
}

func (s *service) GetAvailableNotice(ctx context.Context) (model.Notice, error) {
	return s.repo.GetAvailableNotice(ctx)
}
