package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/pkg/graph/converter"
)

func (r *mutationResolver) Login(ctx context.Context, input converter.LoginInput) (*converter.LoginResp, error) {
	staffInfo, err := r.authSvc.Login(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	resp := &converter.LoginResp{
		StaffID:  staffInfo.ID,
		Username: staffInfo.Username,
		Token:    staffInfo.Token,
	}

	return resp, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	userInfo, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return false, err
	}

	err = r.authSvc.Logout(ctx, userInfo)
	if err != nil {
		return false, err
	}

	return true, nil
}
