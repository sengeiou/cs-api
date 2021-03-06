package faq

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
	authSvc iface.IAuthService
	faqSvc  iface.IFAQService
}

type CreateFAQParams struct {
	Question string       `json:"question" binding:"required"`
	Answer   string       `json:"answer" binding:"required"`
	Status   types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateFAQ(c *gin.Context) {
	var (
		err           error
		requestParams CreateFAQParams
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
	params := model.CreateFAQParams{
		Question:  requestParams.Question,
		Answer:    requestParams.Answer,
		Status:    requestParams.Status,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.faqSvc.CreateFAQ(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateFAQParams struct {
	Question string       `json:"question" binding:"required"`
	Answer   string       `json:"answer" binding:"required"`
	Status   types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateFAQ(c *gin.Context) {
	var (
		err           error
		requestParams UpdateFAQParams
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

	params := model.UpdateFAQParams{
		Question:  requestParams.Question,
		Answer:    requestParams.Answer,
		Status:    requestParams.Status,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        id,
	}

	if err = h.faqSvc.UpdateFAQ(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteFAQ(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.faqSvc.DeleteFAQ(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) ListAvailableFAQ(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)

	tags, err := h.faqSvc.ListAvailableFAQ(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, tags)
}

func (h *handler) GetFAQ(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	tag, err := h.faqSvc.GetFAQ(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, tag)
}

type ListFAQParams struct {
	Question string       `form:"question" binding:""`
	Status   types.Status `form:"status" binding:"min=0,max=2"`
	types.Pagination
}

func (h *handler) ListFAQ(c *gin.Context) {
	var (
		err           error
		requestParams ListFAQParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params, filterParams := formatListFAQParams(requestParams)

	tags, count, err := h.faqSvc.ListFAQ(ctx, params, filterParams)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, tags, requestParams.Pagination)
}

func formatListFAQParams(requestParams ListFAQParams) (model.ListFAQParams, types.FilterFAQParams) {
	params := model.ListFAQParams{
		Limit:  requestParams.Pagination.PageSize,
		Offset: 0,
	}

	filterParams := types.FilterFAQParams{
		Question: nil,
		Status:   nil,
	}

	if requestParams.Question != "" {
		filterParams.Question = &requestParams.Question
	}

	if requestParams.Status != types.StatusAll {
		filterParams.Status = &requestParams.Status
	}

	if requestParams.Pagination.Page > 0 {
		params.Offset = (requestParams.Pagination.Page - 1) * requestParams.Pagination.PageSize
	}

	return params, filterParams
}

func NewHandler(authSvc iface.IAuthService, faqSvc iface.IFAQService) *handler {
	return &handler{
		authSvc: authSvc,
		faqSvc:  faqSvc,
	}
}
