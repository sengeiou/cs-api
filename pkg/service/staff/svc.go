package staff

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	"cs-api/pkg/types"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

func (s *service) ListStaff(ctx context.Context, params model.ListStaffParams, filterParams types.FilterStaffParams) (staffs []model.ListStaffRow, count int64, err error) {
	staffs = make([]model.ListStaffRow, 0)
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

		_, err2 = tx.Exec("SET @servingStatus = ?", filterParams.ServingStatus)
		if err2 != nil {
			return err2
		}

		staffs, err2 = s.repo.WithTx(tx).ListStaff(ctx, params)
		if err2 != nil {
			return err2
		}

		count, err2 = s.repo.WithTx(tx).CountListStaff(ctx)
		if err2 != nil {
			return err2
		}

		return nil
	})

	return
}

func (s *service) GetStaff(ctx context.Context, staffId int64) (staff model.GetStaffRow, err error) {
	return s.repo.GetStaff(ctx, staffId)
}

func (s *service) CreateStaff(ctx context.Context, params model.CreateStaffParams) error {
	return s.repo.CreateStaff(ctx, params)
}

func (s *service) UpdateStaff(ctx context.Context, params interface{}) error {
	switch data := params.(type) {
	case model.UpdateStaffParams:
		return s.repo.UpdateStaff(ctx, data)
	case model.UpdateStaffWithPasswordParams:
		return s.repo.UpdateStaffWithPassword(ctx, data)
	case model.UpdateStaffAvatarParams:
		return s.repo.UpdateStaffAvatar(ctx, data)
	}
	return errors.New("params type error")
}

func (s *service) DeleteStaff(ctx context.Context, staffId int64) error {
	return s.repo.DeleteStaff(ctx, staffId)
}

func (s *service) UpdateStaffServingStatus(ctx context.Context, staffInfo pkg.StaffInfo, status types.StaffServingStatus) error {
	now := time.Now().UTC()

	params := model.UpdateStaffServingStatusParams{
		ServingStatus: status,
		UpdatedBy:     staffInfo.ID,
		UpdatedAt:     now,
		ID:            staffInfo.ID,
	}

	if err := s.repo.UpdateStaffServingStatus(ctx, params); err != nil {
		return err
	}

	mapping := map[types.StaffServingStatus]pkg.StaffEvent{
		types.StaffServingStatusClosed:  pkg.StaffEventClosed,
		types.StaffServingStatusServing: pkg.StaffEventServing,
		types.StaffServingStatusPending: pkg.StaffEventPending,
	}

	event := pkg.StaffEventInfo{
		Event: mapping[status],
		Payload: pkg.StaffEventPayload{
			StaffID: &staffInfo.ID,
		},
	}

	payload, _ := json.Marshal(event)

	newStaffInfo := pkg.StaffInfo{
		ID:            staffInfo.ID,
		Type:          pkg.WsClientTypeStaff,
		Name:          staffInfo.Name,
		Username:      staffInfo.Username,
		ServingStatus: status,
		Token:         staffInfo.Token,
	}

	result, err := json.Marshal(newStaffInfo)
	if err != nil {
		return err
	}

	err = s.redis.SetEX(ctx, "token:staff:"+staffInfo.Token, result, 48*time.Hour)

	// TODO: 後續考慮重構，可以改成用 nats request&reply 方式，如果更新職員服務狀態 event 處理不成功，則切換狀態失敗
	if err = s.redis.Publish(ctx, "event:staff", payload); err != nil {
		return err
	}

	return nil
}

func (s *service) ListAvailableStaff(ctx context.Context, staffId int64) ([]model.Staff, error) {
	return s.repo.ListAvailableStaff(ctx, staffId)
}
