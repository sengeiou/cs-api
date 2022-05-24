package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/pkg"
	"cs-api/pkg/graph/converter"
)

func (r *queryResolver) ListRoomMessage(ctx context.Context, filter converter.ListRoomMessageInput) (*converter.ListRoomMessageResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	messages, err := r.messageSvc.ListRoomMessage(ctx, filter.RoomID, converter.ClientTypeMapping[filter.ClientType])
	if err != nil {
		return nil, err
	}

	resp := &converter.ListRoomMessageResp{}
	resp.FromDto(messages)

	return resp, nil
}

func (r *queryResolver) ListMessage(ctx context.Context, filter converter.ListMessageInput, pagination converter.PaginationInput) (*converter.ListMessageResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	params := pkg.ListMessageParams{
		RoomID:   filter.RoomID,
		StaffID:  filter.StaffID,
		Content:  filter.Content,
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
	}
	messages, total, err := r.messageSvc.ListMessage(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListMessageResp{}
	resp.FromDto(messages, pagination, total)

	return resp, nil
}
