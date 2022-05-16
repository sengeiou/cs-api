package converter

import (
	"cs-api/db/model"
	"cs-api/pkg/types"
	"github.com/golang/go-util/helper"
	"time"
)

var StaffServingStatusMapping = map[StaffServingStatus]types.StaffServingStatus{
	StaffServingStatusAll:     0,
	StaffServingStatusClosed:  types.StaffServingStatusClosed,
	StaffServingStatusServing: types.StaffServingStatusServing,
	StaffServingStatusPending: types.StaffServingStatusPending,
}

var StaffServingStatusDtoMapping = map[types.StaffServingStatus]StaffServingStatus{
	types.StaffServingStatusClosed:  StaffServingStatusClosed,
	types.StaffServingStatusServing: StaffServingStatusServing,
	types.StaffServingStatusPending: StaffServingStatusPending,
}

func (input CreateStaffInput) ConvertToModel(staffId int64, salt string) model.CreateStaffParams {
	now := time.Now().UTC()

	return model.CreateStaffParams{
		RoleID:    input.RoleID,
		Name:      input.Name,
		Username:  input.Username,
		Password:  helper.EncryptPassword(input.Password, salt),
		Status:    StatusMapping[input.Status],
		CreatedBy: staffId,
		CreatedAt: now,
		UpdatedBy: staffId,
		UpdatedAt: now,
	}
}

func (input UpdateStaffInput) ConvertToModel(staffId int64, salt string) interface{} {
	now := time.Now().UTC()

	var params interface{}
	updatePassword := ""
	if input.Password != nil && *input.Password != "" {
		updatePassword = helper.EncryptPassword(*input.Password, salt)
		params = model.UpdateStaffWithPasswordParams{
			RoleID:    input.RoleID,
			Name:      input.Name,
			Password:  updatePassword,
			Status:    StatusMapping[input.Status],
			ID:        input.ID,
			UpdatedBy: staffId,
			UpdatedAt: now,
		}
	} else {
		params = model.UpdateStaffParams{
			RoleID:    input.RoleID,
			Name:      input.Name,
			Status:    StatusMapping[input.Status],
			ID:        input.ID,
			UpdatedBy: staffId,
			UpdatedAt: now,
		}
	}

	return params
}

func (input ListStaffInput) ConvertToModel(pagination PaginationInput) (model.ListStaffParams, types.FilterStaffParams) {
	params := model.ListStaffParams{
		Limit:  int32(pagination.PageSize),
		Offset: 0,
	}

	filterParams := types.FilterStaffParams{
		Name:          nil,
		Status:        nil,
		ServingStatus: nil,
	}

	if input.Name != "" {
		filterParams.Name = &input.Name
	}

	if input.Status != StatusAll {
		status := StatusMapping[input.Status]
		filterParams.Status = &status
	}

	if StaffServingStatusMapping[input.ServingStatus] != 0 {
		servingStatus := StaffServingStatusMapping[input.ServingStatus]
		filterParams.ServingStatus = &servingStatus
	}

	if pagination.Page > 0 {
		params.Offset = int32((pagination.Page - 1) * pagination.PageSize)
	}

	return params, filterParams
}

func (resp *ListStaffResp) FromModel(staffs []model.ListStaffRow, pagination PaginationInput, total int64) {
	resp.Staffs = make([]*Staff, 0)

	for _, staff := range staffs {
		tmp := Staff{
			ID:            staff.ID,
			RoleID:        staff.RoleID,
			Role:          staff.RoleName,
			Name:          staff.Name,
			Username:      staff.Username,
			Status:        StatusModelMapping[staff.Status],
			ServingStatus: StaffServingStatusDtoMapping[staff.ServingStatus],
		}
		resp.Staffs = append(resp.Staffs, &tmp)
	}

	resp.Pagination = &Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    total,
	}
}

func (resp *ListAvailableStaffResp) FromModel(staffs []model.Staff) {
	resp.Staffs = make([]*Staff, 0)

	for _, staff := range staffs {
		tmp := Staff{
			ID:            staff.ID,
			Name:          staff.Name,
			Username:      staff.Username,
			Status:        StatusModelMapping[staff.Status],
			ServingStatus: StaffServingStatusDtoMapping[staff.ServingStatus],
		}
		resp.Staffs = append(resp.Staffs, &tmp)
	}
}
