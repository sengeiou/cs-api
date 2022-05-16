package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg/graph/converter"
	"encoding/json"
	"time"
)

func (r *mutationResolver) CreateStaff(ctx context.Context, input converter.CreateStaffInput) (bool, error) {
	staffInfo, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return false, err
	}

	params := input.ConvertToModel(staffInfo.ID, r.config.Salt)

	err = r.staffSvc.CreateStaff(ctx, params)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateStaff(ctx context.Context, input converter.UpdateStaffInput) (bool, error) {
	staffInfo, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return false, err
	}

	params := input.ConvertToModel(staffInfo.ID, r.config.Salt)

	err = r.staffSvc.UpdateStaff(ctx, params)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteStaff(ctx context.Context, id int64) (bool, error) {
	_, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return false, err
	}

	err = r.staffSvc.DeleteStaff(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateStaffServingStatus(ctx context.Context, servingStatus converter.StaffServingStatus) (bool, error) {
	staffInfo, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return false, err
	}

	err = r.staffSvc.UpdateStaffServingStatus(ctx, staffInfo, converter.StaffServingStatusMapping[servingStatus])
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateStaffAvatar(ctx context.Context, avatar string) (bool, error) {
	staffInfo, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return false, err
	}

	params := model.UpdateStaffAvatarParams{
		Avatar:    avatar,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        staffInfo.ID,
	}

	if err = r.staffSvc.UpdateStaff(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) ListStaff(ctx context.Context, filter converter.ListStaffInput, pagination converter.PaginationInput) (*converter.ListStaffResp, error) {
	_, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return nil, err
	}

	params, filterParams := filter.ConvertToModel(pagination)

	staffs, total, err := r.staffSvc.ListStaff(ctx, params, filterParams)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListStaffResp{}
	resp.FromModel(staffs, pagination, total)

	return resp, nil
}

func (r *queryResolver) ListAvailableStaff(ctx context.Context) (*converter.ListAvailableStaffResp, error) {
	staffInfo, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return nil, err
	}

	staffs, err := r.staffSvc.ListAvailableStaff(ctx, staffInfo.ID)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListAvailableStaffResp{}
	resp.FromModel(staffs)

	return resp, nil
}

func (r *queryResolver) GetStaff(ctx context.Context, id int64) (*converter.GetStaffResp, error) {
	staffInfo, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return nil, err
	}

	if id == 0 {
		id = staffInfo.ID
	}

	staff, err := r.staffSvc.GetStaff(ctx, id)
	if err != nil {
		return nil, err
	}

	tmp := make([]string, 0)
	_ = json.Unmarshal(staff.Permissions, &tmp)

	resp := &converter.GetStaffResp{
		Staff: &converter.Staff{
			ID:            staff.ID,
			RoleID:        staff.RoleID,
			Role:          staff.RoleName,
			Permissions:   tmp,
			Name:          staff.Name,
			Username:      staff.Username,
			Status:        converter.StatusModelMapping[staff.Status],
			ServingStatus: converter.StaffServingStatusDtoMapping[staff.ServingStatus],
			Avatar:        staff.Avatar,
		},
	}

	return resp, nil
}
