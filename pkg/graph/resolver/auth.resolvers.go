package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/pkg"
)

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	userInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	err = r.authSvc.Logout(ctx, userInfo)
	if err != nil {
		return false, err
	}

	return true, nil
}
