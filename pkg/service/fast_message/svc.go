package fast_message

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	"cs-api/pkg/types"
	"database/sql"
)

func (s *service) ListFastMessage(ctx context.Context, params model.ListFastMessageParams, filterParams types.FilterFastMessageParams) (messages []model.ListFastMessageRow, count int64, err error) {
	messages = make([]model.ListFastMessageRow, 0)
	err = s.repo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		var err2 error

		_, err2 = tx.Exec("SET @title = ?", filterParams.Title)
		if err2 != nil {
			return err2
		}

		_, err2 = tx.Exec("SET @content = ?", filterParams.Content)
		if err2 != nil {
			return err2
		}

		_, err2 = tx.Exec("SET @status = ?", filterParams.Status)
		if err2 != nil {
			return err2
		}

		messages, err2 = s.repo.WithTx(tx).ListFastMessage(ctx, params)
		if err2 != nil {
			return err2
		}

		count, err2 = s.repo.WithTx(tx).CountListFastMessage(ctx)
		if err2 != nil {
			return err2
		}

		return nil
	})

	return
}

func (s *service) GetFastMessage(ctx context.Context, id int64) (model.FastMessage, error) {
	return s.repo.GetFastMessage(ctx, id)
}

func (s *service) CreateFastMessage(ctx context.Context, params model.CreateFastMessageParams) error {
	return s.repo.CreateFastMessage(ctx, params)
}

func (s *service) UpdateFastMessage(ctx context.Context, params model.UpdateFastMessageParams) error {
	return s.repo.UpdateFastMessage(ctx, params)
}

func (s *service) DeleteFastMessage(ctx context.Context, id int64) error {
	return s.repo.DeleteFastMessage(ctx, id)
}

func (s *service) ListCategory(ctx context.Context) ([]model.Constant, error) {
	return s.repo.ListFastMessageCategory(ctx)
}

func (s *service) CreateCategory(ctx context.Context, params model.CreateFastMessageCategoryParams) error {
	return s.repo.CreateFastMessageCategory(ctx, params)
}

func (s *service) ListFastMessageGroup(ctx context.Context) ([]pkg.FastMessageGroupItem, error) {
	messages, err := s.repo.GetAllAvailableFastMessage(ctx)
	if err != nil {
		return nil, err
	}

	categoryMap := map[int64]pkg.FastMessageGroupItem{}
	for _, message := range messages {
		if _, ok := categoryMap[message.CategoryID]; !ok {
			tmp := pkg.FastMessageGroupItem{
				Category: pkg.FastMessageCategory{
					ID:   message.CategoryID,
					Name: message.Category,
				},
			}
			tmp.Items = append(tmp.Items, message)
			categoryMap[message.CategoryID] = tmp
		} else {
			tmp := categoryMap[message.CategoryID]
			tmp.Items = append(tmp.Items, message)
			categoryMap[message.CategoryID] = tmp
		}
	}

	group := make([]pkg.FastMessageGroupItem, 0)
	for _, v := range categoryMap {
		group = append(group, v)
	}

	return group, nil
}
