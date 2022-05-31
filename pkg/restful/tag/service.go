package tag

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

func (s *service) ListTag(ctx context.Context, params model.ListTagParams, filterParams types.FilterTagParams) (tags []model.ListTagRow, count int64, err error) {
	tags = make([]model.ListTagRow, 0)
	err = s.repo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		var err2 error

		_, err2 = tx.Exec("SET @name = ?", filterParams.Name)
		if err2 != nil {
			return err2
		}

		_, err2 = tx.Exec("SET @status = ?", filterParams.Status)
		if err2 != nil {
			return err2
		}

		tags, err2 = s.repo.WithTx(tx).ListTag(ctx, params)
		if err2 != nil {
			return err2
		}

		count, err2 = s.repo.WithTx(tx).CountListTag(ctx)
		if err2 != nil {
			return err2
		}

		return nil
	})

	return
}

func (s *service) GetTag(ctx context.Context, tagId int64) (tag model.GetTagRow, err error) {
	return s.repo.GetTag(ctx, tagId)
}

func (s *service) CreateTag(ctx context.Context, params model.CreateTagParams) error {
	return s.repo.CreateTag(ctx, params)
}

func (s *service) UpdateTag(ctx context.Context, params model.UpdateTagParams) error {
	return s.repo.UpdateTag(ctx, params)
}

func (s *service) DeleteTag(ctx context.Context, tagId int64) error {
	return s.repo.DeleteTag(ctx, tagId)
}

func (s *service) ListAvailableTag(ctx context.Context) ([]model.ListAvailableTagRow, error) {
	return s.repo.ListAvailableTag(ctx)
}

func NewService(Repo iface.IRepository) iface.ITagService {
	return &service{
		repo: Repo,
	}
}
