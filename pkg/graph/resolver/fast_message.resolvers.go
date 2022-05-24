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

func (r *mutationResolver) CreateFastMessage(ctx context.Context, input converter.CreateFastMessageInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	now := time.Now().UTC()

	params := model.CreateFastMessageParams{
		CategoryID: input.CategoryID,
		Title:      input.Title,
		Content:    input.Content,
		Status:     converter.StatusMapping[input.Status],
		CreatedBy:  staffInfo.ID,
		CreatedAt:  now,
		UpdatedBy:  staffInfo.ID,
		UpdatedAt:  now,
	}

	if err = r.fastMessageSvc.CreateFastMessage(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateFastMessage(ctx context.Context, input converter.UpdateFastMessageInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	params := model.UpdateFastMessageParams{
		CategoryID: input.CategoryID,
		Title:      input.Title,
		Content:    input.Content,
		Status:     converter.StatusMapping[input.Status],
		UpdatedBy:  staffInfo.ID,
		UpdatedAt:  time.Now().UTC(),
		ID:         staffInfo.ID,
	}

	if err = r.fastMessageSvc.UpdateFastMessage(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteFastMessage(ctx context.Context, id int64) (bool, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	if err = r.fastMessageSvc.DeleteFastMessage(ctx, id); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) CreateFastMessageCategory(ctx context.Context, input converter.CreateFastMessageCategoryInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	now := time.Now().UTC()

	params := model.CreateFastMessageCategoryParams{
		Value:     input.Name,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = r.fastMessageSvc.CreateCategory(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) ListFastMessage(ctx context.Context, filter converter.ListFastMessageInput, pagination converter.PaginationInput) (*converter.ListFastMessageResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	params, filterParams := filter.ConvertToModel(pagination)

	messages, total, err := r.fastMessageSvc.ListFastMessage(ctx, params, filterParams)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListFastMessageResp{}
	resp.FromModel(messages, pagination, total)

	return resp, nil
}

func (r *queryResolver) ListFastMessageCategory(ctx context.Context) (*converter.ListFastMessageCategoryResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	categories, err := r.fastMessageSvc.ListCategory(ctx)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListFastMessageCategoryResp{}
	resp.FromModel(categories)

	return resp, nil
}

func (r *queryResolver) ListFastMessageGroup(ctx context.Context) (*converter.ListFastMessageGroupResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	group, err := r.fastMessageSvc.ListFastMessageGroup(ctx)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListFastMessageGroupResp{}
	resp.FromModel(group)

	return resp, nil
}

func (r *queryResolver) GetFastMessage(ctx context.Context, id int64) (*converter.GetFastMessageResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	message, err := r.fastMessageSvc.GetFastMessage(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &converter.GetFastMessageResp{
		FastMessage: &converter.FastMessage{
			ID:         message.ID,
			CategoryID: message.CategoryID,
			Title:      message.Title,
			Content:    message.Content,
			Status:     converter.StatusModelMapping[message.Status],
		},
	}

	return resp, nil
}
