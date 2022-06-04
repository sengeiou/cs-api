package message

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
	"errors"
)

type service struct {
	repo iface.IRepository
}

func (s *service) CreateMessage(ctx context.Context, params model.CreateMessageParams) error {
	return s.repo.CreateMessage(ctx, params)
}

func (s *service) ListRoomMessage(ctx context.Context, params interface{}) (messages []model.Message, err error) {
	messages = make([]model.Message, 0)

	switch data := params.(type) {
	case model.ListMemberRoomMessageParams:
		return s.repo.ListMemberRoomMessage(ctx, data)
	case model.ListStaffRoomMessageParams:
		return s.repo.ListStaffRoomMessage(ctx, data)
	}

	return make([]model.Message, 0), errors.New("wrong client type")
}

func (s *service) ListMessage(ctx context.Context, params model.ListMessageParams, filterParams types.FilterMessageParams) (messages []model.Message, count int64, err error) {
	messages = make([]model.Message, 0)

	err = s.repo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		var err2 error

		_, err2 = tx.Exec("SET @roomId = ?", filterParams.RoomID)
		if err2 != nil {
			return err2
		}

		_, err2 = tx.Exec("SET @staffId = ?", filterParams.StaffID)
		if err2 != nil {
			return err2
		}

		_, err2 = tx.Exec("SET @content = ?", filterParams.Content)
		if err2 != nil {
			return err2
		}

		messages, err2 = s.repo.WithTx(tx).ListMessage(ctx, params)
		if err2 != nil {
			return err2
		}

		count, err2 = s.repo.WithTx(tx).CountListMessage(ctx)
		if err2 != nil {
			return err2
		}

		return nil
	})

	return
}

func NewService(repo iface.IRepository) iface.IMessageService {
	return &service{
		repo: repo,
	}
}
