package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/pkg"
	"cs-api/pkg/graph/converter"
	"cs-api/pkg/types"
)

func (r *mutationResolver) UpdateCsConfig(ctx context.Context, input converter.UpdateCsConfigInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	config := types.CsConfig{
		MaxMember:           input.MaxMember,
		MemberPendingExpire: input.MemberPendingExpire,
		GreetingText:        input.GreetingText,
	}

	if err = r.csConfigSvc.UpdateCsConfig(ctx, staffInfo.ID, config); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) GetCsConfig(ctx context.Context) (*converter.GetCsConfigResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	config, err := r.csConfigSvc.GetCsConfig(ctx)
	if err != nil {
		return nil, err
	}

	resp := &converter.GetCsConfigResp{
		Config: &converter.CsConfig{
			MaxMember:           config.MaxMember,
			MemberPendingExpire: config.MemberPendingExpire,
			GreetingText:        config.GreetingText,
		},
	}

	return resp, nil
}
