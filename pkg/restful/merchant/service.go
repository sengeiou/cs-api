package merchant

import (
	"context"
	"crypto/md5"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
)

type service struct {
	repo iface.IRepository
}

func (s *service) ListMerchant(ctx context.Context, params model.ListMerchantParams, filterParams types.FilterMerchantParams) (merchants []model.ListMerchantRow, count int64, err error) {
	merchants = make([]model.ListMerchantRow, 0)
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

		merchants, err2 = s.repo.WithTx(tx).ListMerchant(ctx, params)
		if err2 != nil {
			return err2
		}

		count, err2 = s.repo.WithTx(tx).CountListMerchant(ctx)
		if err2 != nil {
			return err2
		}

		return nil
	})

	return
}

func (s *service) GetMerchant(ctx context.Context, merchantId int64) (merchant model.GetMerchantRow, err error) {
	return s.repo.GetMerchant(ctx, merchantId)
}

func (s *service) CreateMerchant(ctx context.Context, params model.CreateMerchantParams) error {
	name, code := params.Name, params.Code
	key := fmt.Sprintf("%x", md5.Sum([]byte(name+code+strconv.Itoa(rand.Intn(999)))))[:16]
	for {
		id, err := s.repo.CheckMerchantKey(ctx, key)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if id != 0 {
			key = fmt.Sprintf("%x", md5.Sum([]byte(name+code+strconv.Itoa(rand.Intn(999)))))[:16]
		} else {
			params.Key = key
			break
		}
	}
	return s.repo.CreateMerchant(ctx, params)
}

func (s *service) UpdateMerchant(ctx context.Context, params model.UpdateMerchantParams) error {
	return s.repo.UpdateMerchant(ctx, params)
}

func (s *service) DeleteMerchant(ctx context.Context, merchantId int64) error {
	return s.repo.DeleteMerchant(ctx, merchantId)
}

func (s *service) ListAvailableMerchant(ctx context.Context) ([]model.ListAvailableMerchantRow, error) {
	return s.repo.ListAvailableMerchant(ctx)
}

func NewService(repo iface.IRepository) iface.IMerchantService {
	return &service{
		repo: repo,
	}
}
