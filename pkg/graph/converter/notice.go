package converter

import (
	"cs-api/db/model"
	"cs-api/pkg/types"
	"time"
)

func (input CreateNoticeInput) ConvertToModel(staffId int64) model.CreateNoticeParams {
	now := time.Now().UTC()

	startTime, _ := time.Parse("2006-01-02 15:04:05", input.StartAt)
	startTime = startTime.Add(-8 * time.Hour)

	endTime, _ := time.Parse("2006-01-02 15:04:05", input.EndAt)
	endTime = endTime.Add(-8 * time.Hour)

	return model.CreateNoticeParams{
		Title:     input.Title,
		Content:   input.Content,
		StartAt:   startTime,
		EndAt:     endTime,
		Status:    StatusMapping[input.Status],
		CreatedBy: staffId,
		CreatedAt: now,
		UpdatedBy: staffId,
		UpdatedAt: now,
	}
}

func (input UpdateNoticeInput) ConvertToModel(staffId int64) model.UpdateNoticeParams {
	startTime, _ := time.Parse("2006-01-02 15:04:05", input.StartAt)
	startTime = startTime.Add(-8 * time.Hour)

	endTime, _ := time.Parse("2006-01-02 15:04:05", input.EndAt)
	endTime = endTime.Add(-8 * time.Hour)

	return model.UpdateNoticeParams{
		Title:     input.Title,
		Content:   input.Content,
		StartAt:   startTime,
		EndAt:     endTime,
		Status:    StatusMapping[input.Status],
		UpdatedBy: staffId,
		UpdatedAt: time.Now().UTC(),
		ID:        input.ID,
	}
}

func (input ListNoticeInput) ConvertToModel(pagination PaginationInput) (model.ListNoticeParams, types.FilterNoticeParams) {
	params := model.ListNoticeParams{
		Limit:  int32(pagination.PageSize),
		Offset: 0,
	}

	filterParams := types.FilterNoticeParams{
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

func (resp *ListNoticeResp) FromModel(notices []model.Notice, pagination PaginationInput, total int64) {
	resp.Notices = make([]*Notice, 0, len(notices))

	for _, notice := range notices {
		tmp := Notice{
			ID:      notice.ID,
			Title:   notice.Title,
			Content: notice.Content,
			StartAt: FormatTime(notice.StartAt),
			EndAt:   FormatTime(notice.EndAt),
			Status:  StatusModelMapping[notice.Status],
		}
		resp.Notices = append(resp.Notices, &tmp)
	}

	resp.Pagination = &Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    total,
	}
}

func (resp *GetNoticeResp) FromModel(notice model.Notice) {
	resp.Notice = &Notice{
		ID:      notice.ID,
		Title:   notice.Title,
		Content: notice.Content,
		StartAt: FormatTime(notice.StartAt),
		EndAt:   FormatTime(notice.EndAt),
		Status:  StatusModelMapping[notice.Status],
	}
}
