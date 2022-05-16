package converter

import (
	"cs-api/db/model"
	"cs-api/pkg/types"
	"time"
)

func (input CreateTagInput) ConvertToModel(staffId int64) model.CreateTagParams {
	now := time.Now().UTC()

	return model.CreateTagParams{
		Name:      input.Name,
		Status:    StatusMapping[input.Status],
		CreatedBy: staffId,
		CreatedAt: now,
		UpdatedBy: staffId,
		UpdatedAt: now,
	}
}

func (input ListTagInput) ConvertToModel(pagination PaginationInput) (model.ListTagParams, types.FilterTagParams) {
	params := model.ListTagParams{
		Limit:  int32(pagination.PageSize),
		Offset: 0,
	}

	filterParams := types.FilterTagParams{
		Name:   nil,
		Status: nil,
	}

	if input.Name != "" {
		filterParams.Name = &input.Name
	}

	if input.Status != StatusAll {
		status := StatusMapping[input.Status]
		filterParams.Status = &status
	}

	if pagination.Page > 0 {
		params.Offset = int32((pagination.Page - 1) * pagination.PageSize)
	}

	return params, filterParams
}

func (resp *ListTagResp) FromModel(tags []model.Tag, pagination PaginationInput, total int64) {
	resp.Tags = make([]*Tag, 0)

	for _, tag := range tags {
		tmp := Tag{
			ID:     tag.ID,
			Name:   tag.Name,
			Status: StatusModelMapping[tag.Status],
		}
		resp.Tags = append(resp.Tags, &tmp)
	}

	resp.Pagination = &Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    total,
	}
}
