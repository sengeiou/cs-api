package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/pkg"
	"cs-api/pkg/graph/converter"
	"encoding/json"
)

func (r *mutationResolver) CreateRole(ctx context.Context, input converter.CreateRoleInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	params := input.ConvertToModel(staffInfo.ID)

	if err = r.roleSvc.CreateRole(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateRole(ctx context.Context, input converter.UpdateRoleInput) (bool, error) {
	staffInfo, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	params := input.ConvertToModel(staffInfo.ID)

	if err = r.roleSvc.UpdateRole(ctx, params); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteRole(ctx context.Context, id int64) (bool, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return false, err
	}

	if err = r.roleSvc.DeleteRole(ctx, id); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) ListRole(ctx context.Context, filter converter.ListRoleInput, pagination converter.PaginationInput) (*converter.ListRoleResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	params, filterParams := filter.ConvertToModel(pagination)

	roles, total, err := r.roleSvc.ListRole(ctx, params, filterParams)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListRoleResp{}
	if err = resp.FromModel(roles, pagination, total); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *queryResolver) GetRole(ctx context.Context, id int64) (*converter.GetRoleResp, error) {
	_, err := r.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		return nil, err
	}

	role, err := r.roleSvc.GetRole(ctx, id)
	if err != nil {
		return nil, err
	}

	tmp := make([]string, 0)
	if err = json.Unmarshal(role.Permissions, &tmp); err != nil {
		return nil, err
	}

	resp := &converter.GetRoleResp{
		Role: &converter.Role{
			ID:          role.ID,
			Name:        role.Name,
			Permissions: tmp,
		},
	}

	return resp, nil
}
