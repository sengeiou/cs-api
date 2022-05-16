package converter

import (
	"cs-api/db/model"
	"cs-api/pkg"
	"cs-api/pkg/types"
)

func (input ListFastMessageInput) ConvertToModel(pagination PaginationInput) (model.ListFastMessageParams, types.FilterFastMessageParams) {
	params := model.ListFastMessageParams{
		Limit:  int32(pagination.PageSize),
		Offset: 0,
	}

	filterParams := types.FilterFastMessageParams{
		Title:   nil,
		Content: nil,
		Status:  nil,
	}

	if input.Title != "" {
		filterParams.Title = &input.Title
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

func (resp *ListFastMessageResp) FromModel(messages []model.ListFastMessageRow, pagination PaginationInput, total int64) {
	resp.FastMessages = make([]*FastMessage, 0)

	for _, message := range messages {
		tmp := FastMessage{
			ID:         message.ID,
			Category:   message.Category,
			CategoryID: message.CategoryID,
			Title:      message.Title,
			Content:    message.Content,
			Status:     StatusModelMapping[message.Status],
		}
		resp.FastMessages = append(resp.FastMessages, &tmp)
	}

	resp.Pagination = &Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    total,
	}
}

func (resp *ListFastMessageCategoryResp) FromModel(categories []model.Constant) {
	resp.Categories = make([]*FastMessageCategory, 0)

	for _, category := range categories {
		tmp := FastMessageCategory{
			ID:   category.ID,
			Name: category.Value,
		}
		resp.Categories = append(resp.Categories, &tmp)
	}
}

func (resp *ListFastMessageGroupResp) FromModel(group []pkg.FastMessageGroupItem) {
	resp.Group = make([]*FastMessageGroupItem, 0, len(group))

	for _, groupItem := range group {
		tmp := FastMessageGroupItem{
			Category: &FastMessageCategory{
				ID:   groupItem.Category.ID,
				Name: groupItem.Category.Name,
			},
		}

		for _, item := range groupItem.Items {
			tmp.Items = append(tmp.Items, &FastMessage{
				ID:         item.ID,
				Category:   item.Category,
				CategoryID: item.CategoryID,
				Title:      item.Title,
				Content:    item.Content,
				Status:     StatusModelMapping[item.Status],
			})
		}

		resp.Group = append(resp.Group, &tmp)
	}
}
