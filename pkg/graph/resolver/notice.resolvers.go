package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/pkg"
	"cs-api/pkg/graph/converter"
)

func (r *mutationResolver) CreateNotice(ctx context.Context, input converter.CreateNoticeInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	params := input.ConvertToModel(staffInfo.ID)

	if err = r.noticeSvc.CreateNotice(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateNotice(ctx context.Context, input converter.UpdateNoticeInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	params := input.ConvertToModel(staffInfo.ID)

	if err = r.noticeSvc.UpdateNotice(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteNotice(ctx context.Context, id int64) (bool, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	if err = r.noticeSvc.DeleteNotice(ctx, id); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) ListNotice(ctx context.Context, filter converter.ListNoticeInput, pagination converter.PaginationInput) (*converter.ListNoticeResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	params, filterParams := filter.ConvertToModel(pagination)

	notices, total, err := r.noticeSvc.ListNotice(ctx, params, filterParams)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListNoticeResp{}
	resp.FromModel(notices, pagination, total)

	return resp, nil
}

func (r *queryResolver) GetNotice(ctx context.Context, id int64) (*converter.GetNoticeResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	notice, err := r.noticeSvc.GetNotice(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &converter.GetNoticeResp{}
	resp.FromModel(notice)

	return resp, nil
}
