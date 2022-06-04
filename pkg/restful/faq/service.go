package faq

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

func (s *service) ListFAQ(ctx context.Context, params model.ListFAQParams, filterParams types.FilterFAQParams) (FAQs []model.ListFAQRow, count int64, err error) {
	FAQs = make([]model.ListFAQRow, 0)
	err = s.repo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		var err2 error

		_, err2 = tx.Exec("SET @question = ?", filterParams.Question)
		if err2 != nil {
			return err2
		}

		_, err2 = tx.Exec("SET @status = ?", filterParams.Status)
		if err2 != nil {
			return err2
		}

		FAQs, err2 = s.repo.WithTx(tx).ListFAQ(ctx, params)
		if err2 != nil {
			return err2
		}

		count, err2 = s.repo.WithTx(tx).CountListFAQ(ctx)
		if err2 != nil {
			return err2
		}

		return nil
	})

	return
}

func (s *service) GetFAQ(ctx context.Context, faqId int64) (model.GetFAQRow, error) {
	return s.repo.GetFAQ(ctx, faqId)
}

func (s *service) CreateFAQ(ctx context.Context, params model.CreateFAQParams) error {
	return s.repo.CreateFAQ(ctx, params)
}

func (s *service) UpdateFAQ(ctx context.Context, params model.UpdateFAQParams) error {
	return s.repo.UpdateFAQ(ctx, params)
}

func (s *service) DeleteFAQ(ctx context.Context, faqId int64) error {
	return s.repo.DeleteFAQ(ctx, faqId)
}

func (s *service) ListAvailableFAQ(ctx context.Context) ([]model.ListAvailableFAQRow, error) {
	return s.repo.ListAvailableFAQ(ctx)
}

func NewService(Repo iface.IRepository) iface.IFAQService {
	return &service{
		repo: Repo,
	}
}
