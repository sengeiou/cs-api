package converter

import (
	"cs-api/db/model"
	"cs-api/pkg/types"
	"time"
)

func (input CreateRemindInput) ConvertToModel(staffId int64) model.CreateRemindParams {
	now := time.Now().UTC()

	return model.CreateRemindParams{
		Title:     input.Title,
		Content:   input.Content,
		Status:    StatusMapping[input.Status],
		CreatedBy: staffId,
		CreatedAt: now,
		UpdatedBy: staffId,
		UpdatedAt: now,
	}
}

func (input UpdateRemindInput) ConvertToModel(staffId int64) model.UpdateRemindParams {
	return model.UpdateRemindParams{
		Title:     input.Title,
		Content:   input.Content,
		Status:    StatusMapping[input.Status],
		UpdatedBy: staffId,
		UpdatedAt: time.Now().UTC(),
		ID:        input.ID,
	}
}

func (input ListRemindInput) ConvertToModel(pagination PaginationInput) (model.ListRemindParams, types.FilterRemindParams) {
	params := model.ListRemindParams{
		Limit:  int32(pagination.PageSize),
		Offset: 0,
	}

	filterParams := types.FilterRemindParams{
		Content: nil,
		Status:  nil,
	}

	if input.Content != "" {
		filterParams.Content = &input.Content
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

func (resp *ListRemindResp) FromModel(reminds []model.Remind, pagination PaginationInput, total int64) {
	resp.Reminds = make([]*Remind, 0, len(reminds))

	for _, remind := range reminds {
		tmp := Remind{
			ID:      remind.ID,
			Title:   remind.Title,
			Content: remind.Content,
			Status:  StatusModelMapping[remind.Status],
		}
		resp.Reminds = append(resp.Reminds, &tmp)
	}

	resp.Pagination = &Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    total,
	}
}

func (resp *GetRemindResp) FromModel(remind model.Remind) {
	resp.Remind = &Remind{
		ID:      remind.ID,
		Title:   remind.Title,
		Content: remind.Content,
		Status:  StatusModelMapping[remind.Status],
	}
}
