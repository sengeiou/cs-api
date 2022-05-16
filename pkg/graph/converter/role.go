package converter

import (
	"cs-api/db/model"
	"cs-api/pkg/types"
	"encoding/json"
	"time"
)

func (input CreateRoleInput) ConvertToModel(staffId int64) model.CreateRoleParams {
	now := time.Now().UTC()

	result, _ := json.Marshal(input.Permissions)

	return model.CreateRoleParams{
		Name:        input.Name,
		Permissions: result,
		CreatedBy:   staffId,
		CreatedAt:   now,
		UpdatedBy:   staffId,
		UpdatedAt:   now,
	}
}

func (input UpdateRoleInput) ConvertToModel(staffId int64) model.UpdateRoleParams {
	now := time.Now().UTC()

	result, _ := json.Marshal(input.Permissions)

	return model.UpdateRoleParams{
		Name:        input.Name,
		Permissions: result,
		UpdatedBy:   staffId,
		UpdatedAt:   now,
		ID:          input.ID,
	}
}

func (input ListRoleInput) ConvertToModel(pagination PaginationInput) (model.ListRoleParams, types.FilterRoleParams) {
	params := model.ListRoleParams{
		Limit:  int32(pagination.PageSize),
		Offset: 0,
	}

	filterParams := types.FilterRoleParams{
		Name: nil,
	}

	if input.Name != "" {
		filterParams.Name = &input.Name
	}

	if pagination.Page > 0 {
		params.Offset = int32((pagination.Page - 1) * pagination.PageSize)
	}

	return params, filterParams
}

func (resp *ListRoleResp) FromModel(tags []model.Role, pagination PaginationInput, total int64) error {
	resp.Roles = make([]*Role, 0)

	for _, tag := range tags {
		tmpResult := make([]string, 0)
		if err := json.Unmarshal(tag.Permissions, &tmpResult); err != nil {
			return err
		}
		tmp := Role{
			ID:          tag.ID,
			Name:        tag.Name,
			Permissions: tmpResult,
		}
		resp.Roles = append(resp.Roles, &tmp)
	}

	resp.Pagination = &Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    total,
	}

	return nil
}
