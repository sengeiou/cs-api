package notice

import (
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type handler struct {
	authSvc   iface.IAuthService
	noticeSvc iface.INoticeService
}

type CreateNoticeParams struct {
	Title   string         `json:"title" binding:"required"`
	Content string         `json:"content" binding:"required"`
	StartAt types.JSONTime `json:"start_at" binding:"required"`
	EndAt   types.JSONTime `json:"end_at" binding:"required"`
	Status  types.Status   `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateNotice(c *gin.Context) {
	var (
		err           error
		requestParams CreateNoticeParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	now := time.Now().UTC()

	params := model.CreateNoticeParams{
		Title:     requestParams.Title,
		Content:   requestParams.Content,
		StartAt:   requestParams.StartAt.Time,
		EndAt:     requestParams.EndAt.Time,
		Status:    requestParams.Status,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.noticeSvc.CreateNotice(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateNoticeParams struct {
	Title   string         `json:"title" binding:"required"`
	Content string         `json:"content" binding:"required"`
	StartAt types.JSONTime `json:"start_at" binding:"required"`
	EndAt   types.JSONTime `json:"end_at" binding:"required"`
	Status  types.Status   `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateNotice(c *gin.Context) {
	var (
		err           error
		requestParams UpdateNoticeParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params := model.UpdateNoticeParams{
		Title:     requestParams.Title,
		Content:   requestParams.Content,
		StartAt:   requestParams.StartAt.Time,
		EndAt:     requestParams.EndAt.Time,
		Status:    requestParams.Status,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        id,
	}

	if err = h.noticeSvc.UpdateNotice(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteNotice(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.noticeSvc.DeleteNotice(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) GetNotice(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	notice, err := h.noticeSvc.GetNotice(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, notice)
}

type ListNoticeParams struct {
	Content string       `form:"content" binding:""`
	Status  types.Status `form:"status" binding:"min=0,max=2"`
	types.Pagination
}

func (h *handler) ListNotice(c *gin.Context) {
	var (
		err           error
		requestParams ListNoticeParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params, filterParams := formatListNoticeParams(requestParams)

	notices, count, err := h.noticeSvc.ListNotice(ctx, params, filterParams)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, notices, requestParams.Pagination)
}

func formatListNoticeParams(requestParams ListNoticeParams) (model.ListNoticeParams, types.FilterNoticeParams) {
	params := model.ListNoticeParams{
		Limit:  requestParams.PageSize,
		Offset: 0,
	}

	filterParams := types.FilterNoticeParams{
		Content: nil,
		Status:  nil,
	}

	if requestParams.Content != "" {
		filterParams.Content = &requestParams.Content
	}

	if requestParams.Status != types.StatusAll {
		filterParams.Status = &requestParams.Status
	}

	if requestParams.Page > 0 {
		params.Offset = (requestParams.Page - 1) * requestParams.PageSize
	}

	return params, filterParams
}

func (h *handler) GetLatestNotice(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	notice, err := h.noticeSvc.GetLatestNotice(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ginTool.Success(c)
			return
		}
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, notice)
}

func NewHandler(authSvc iface.IAuthService, noticeSvc iface.INoticeService) *handler {
	return &handler{
		authSvc:   authSvc,
		noticeSvc: noticeSvc,
	}
}
