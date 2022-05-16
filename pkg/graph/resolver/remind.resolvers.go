package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/pkg/graph/converter"
)

func (r *mutationResolver) CreateRemind(ctx context.Context, input converter.CreateRemindInput) (bool, error) {
	staffInfo, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return false, err
	}

	params := input.ConvertToModel(staffInfo.ID)

	if err = r.remindSvc.CreateRemind(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateRemind(ctx context.Context, input converter.UpdateRemindInput) (bool, error) {
	staffInfo, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return false, err
	}

	params := input.ConvertToModel(staffInfo.ID)

	if err = r.remindSvc.UpdateRemind(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteRemind(ctx context.Context, id int64) (bool, error) {
	_, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return false, err
	}

	if err = r.remindSvc.DeleteRemind(ctx, id); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) ListRemind(ctx context.Context, filter converter.ListRemindInput, pagination converter.PaginationInput) (*converter.ListRemindResp, error) {
	_, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return nil, err
	}

	params, filterParams := filter.ConvertToModel(pagination)

	notices, total, err := r.remindSvc.ListRemind(ctx, params, filterParams)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListRemindResp{}
	resp.FromModel(notices, pagination, total)

	return resp, nil
}

func (r *queryResolver) GetRemind(ctx context.Context, id int64) (*converter.GetRemindResp, error) {
	_, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return nil, err
	}

	remind, err := r.remindSvc.GetRemind(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &converter.GetRemindResp{}
	resp.FromModel(remind)

	return resp, nil
}
