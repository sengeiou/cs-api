package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	"cs-api/pkg/graph/converter"
	"time"
)

func (r *mutationResolver) CreateTag(ctx context.Context, input converter.CreateTagInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	params := input.ConvertToModel(staffInfo.ID)

	if err = r.tagSvc.CreateTag(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateTag(ctx context.Context, input converter.UpdateTagInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	params := model.UpdateTagParams{
		Name:      input.Name,
		Status:    converter.StatusMapping[input.Status],
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        input.ID,
	}

	if err = r.tagSvc.UpdateTag(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id int64) (bool, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	if err = r.tagSvc.DeleteTag(ctx, id); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) ListTag(ctx context.Context, filter converter.ListTagInput, pagination converter.PaginationInput) (*converter.ListTagResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	params, filterParams := filter.ConvertToModel(pagination)

	tags, total, err := r.tagSvc.ListTag(ctx, params, filterParams)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListTagResp{}
	resp.FromModel(tags, pagination, total)

	return resp, nil
}

func (r *queryResolver) GetTag(ctx context.Context, id int64) (*converter.GetTagResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	tag, err := r.tagSvc.GetTag(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &converter.GetTagResp{
		Tag: &converter.Tag{
			ID:     tag.ID,
			Name:   tag.Name,
			Status: converter.StatusModelMapping[tag.Status],
		},
	}

	return resp, nil
}
