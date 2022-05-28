package fast_reply

import (
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type handler struct {
	authSvc      iface.IAuthService
	fastReplySvc iface.IFastReplyService
}

type CreateFastReplyParams struct {
	CategoryID int64        `json:"category_id" binding:"required"`
	Title      string       `json:"title" binding:"required"`
	Content    string       `json:"content" binding:"required"`
	Status     types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateFastReply(c *gin.Context) {
	var (
		err           error
		requestParams CreateFastReplyParams
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
	params := model.CreateFastReplyParams{
		CategoryID: requestParams.CategoryID,
		Title:      requestParams.Title,
		Content:    requestParams.Content,
		Status:     requestParams.Status,
		CreatedBy:  staffInfo.ID,
		CreatedAt:  now,
		UpdatedBy:  staffInfo.ID,
		UpdatedAt:  now,
	}

	_, err = h.fastReplySvc.CheckCategory(ctx, requestParams.CategoryID)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.fastReplySvc.CreateFastReply(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateFastReplyParams struct {
	CategoryID int64        `json:"category_id" binding:"required"`
	Title      string       `json:"title" binding:"required"`
	Content    string       `json:"content" binding:"required"`
	Status     types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateFastReply(c *gin.Context) {
	var (
		err           error
		requestParams UpdateFastReplyParams
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

	params := model.UpdateFastReplyParams{
		CategoryID: requestParams.CategoryID,
		Title:      requestParams.Title,
		Content:    requestParams.Content,
		Status:     requestParams.Status,
		UpdatedBy:  staffInfo.ID,
		UpdatedAt:  time.Now().UTC(),
		ID:         id,
	}

	if err = h.fastReplySvc.UpdateFastReply(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteFastReply(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.fastReplySvc.DeleteFastReply(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) GetFastReply(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	fastReply, err := h.fastReplySvc.GetFastReply(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, fastReply)
}

type ListFastReplyParams struct {
	Title   string       `form:"title" binding:""`
	Content string       `form:"content" binding:""`
	Status  types.Status `form:"status" binding:"min=0,max=2"`
	types.Pagination
}

func (h *handler) ListFastReply(c *gin.Context) {
	var (
		err           error
		requestParams ListFastReplyParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params, filterParams := formatListFastReplyParams(requestParams)

	fastReplies, count, err := h.fastReplySvc.ListFastReply(ctx, params, filterParams)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, fastReplies, requestParams.Pagination)
}

func formatListFastReplyParams(requestParams ListFastReplyParams) (model.ListFastReplyParams, types.FilterFastReplyParams) {
	params := model.ListFastReplyParams{
		Limit:  requestParams.PageSize,
		Offset: 0,
	}

	filterParams := types.FilterFastReplyParams{
		Title:   nil,
		Content: nil,
		Status:  nil,
	}

	if requestParams.Title != "" {
		filterParams.Title = &requestParams.Title
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

func (h *handler) ListFastReplyGroup(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	group, err := h.fastReplySvc.ListFastReplyGroup(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, group)
}

func (h *handler) ListFastReplyCategory(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	categories, err := h.fastReplySvc.ListCategory(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, categories)
}

type CreateFastReplyCategoryParams struct {
	Name string `json:"name" binding:"required"`
}

func (h *handler) CreateFastReplyCategory(c *gin.Context) {
	var (
		err           error
		requestParams CreateFastReplyCategoryParams
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
	params := model.CreateFastReplyCategoryParams{
		Value:     requestParams.Name,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.fastReplySvc.CreateCategory(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func NewHandler(authSvc iface.IAuthService, fastReplySvc iface.IFastReplyService) *handler {
	return &handler{
		authSvc:      authSvc,
		fastReplySvc: fastReplySvc,
	}
}
