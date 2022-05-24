package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/pkg"
	"cs-api/pkg/graph/converter"
	"errors"
)

func (r *mutationResolver) AcceptRoom(ctx context.Context, id int64) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	if err = r.roomSvc.AcceptRoom(ctx, staffInfo.ID, id); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) CloseRoom(ctx context.Context, input converter.CloseRoomInput) (bool, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	if err = r.roomSvc.CloseRoom(ctx, input.ID, input.TagID); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) TransferRoom(ctx context.Context, input converter.TransferRoomInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	if staffInfo.ID == input.StaffID {
		return false, errors.New("no need to transfer")
	}

	if err = r.roomSvc.TransferRoom(ctx, input.ID, input.StaffID); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateRoomScore(ctx context.Context, score int64) (bool, error) {
	if score < 1 || score > 5 {
		return false, errors.New("wrong score")
	}

	memberInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeMember)
	if err != nil {
		return false, err
	}

	if err = r.roomSvc.UpdateRoomScore(ctx, memberInfo.RoomID, int32(score)); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) ListStaffRoom(ctx context.Context, filter converter.ListStaffRoomInput, pagination converter.PaginationInput) (*converter.ListStaffRoomResp, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	params, filterParams := filter.ConvertToModel(staffInfo.ID, pagination)

	rooms, total, err := r.roomSvc.ListStaffRoom(ctx, params, filterParams)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListStaffRoomResp{}
	resp.FromModel(rooms, pagination, total)

	return resp, nil
}

func (r *queryResolver) ListRoom(ctx context.Context, filter converter.ListRoomInput, pagination converter.PaginationInput) (*converter.ListRoomResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	params, filterParams := filter.ConvertToModel(pagination)

	rooms, total, err := r.roomSvc.ListRoom(ctx, params, filterParams)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListRoomResp{}
	resp.FromModel(rooms, pagination, total)

	return resp, nil
}
