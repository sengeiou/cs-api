package tag

import (
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"fmt"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type handler struct {
	authSvc iface.IAuthService
	tagSvc  iface.ITagService
}

type CreateTagParams struct {
	Name   string       `json:"name" binding:"required"`
	Status types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateTag(c *gin.Context) {
	var (
		err           error
		requestParams CreateTagParams
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
	params := model.CreateTagParams{
		Name:      requestParams.Name,
		Status:    requestParams.Status,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.tagSvc.CreateTag(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateTagParams struct {
	Name   string       `json:"name" binding:"required"`
	Status types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateTag(c *gin.Context) {
	var (
		err           error
		requestParams UpdateTagParams
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

	params := model.UpdateTagParams{
		Name:      requestParams.Name,
		Status:    requestParams.Status,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        id,
	}

	if err = h.tagSvc.UpdateTag(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteTag(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.tagSvc.DeleteTag(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) GetTag(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	tag, err := h.tagSvc.GetTag(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, tag)
}

type ListTagParams struct {
	Name   string       `form:"name" binding:""`
	Status types.Status `form:"status" binding:"required"`
	types.Pagination
}

func (h *handler) ListTag(c *gin.Context) {
	var (
		err           error
		requestParams ListTagParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		fmt.Println(err)
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params, filterParams := formatListTagParams(requestParams)

	tags, count, err := h.tagSvc.ListTag(ctx, params, filterParams)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, tags, requestParams.Pagination)
}

func formatListTagParams(requestParams ListTagParams) (model.ListTagParams, types.FilterTagParams) {
	params := model.ListTagParams{
		Limit:  requestParams.Pagination.PageSize,
		Offset: 0,
	}

	filterParams := types.FilterTagParams{
		Name:   nil,
		Status: nil,
	}

	if requestParams.Name != "" {
		filterParams.Name = &requestParams.Name
	}

	if requestParams.Status != types.StatusAll {
		filterParams.Status = &requestParams.Status
	}

	if requestParams.Pagination.Page > 0 {
		params.Offset = (requestParams.Pagination.Page - 1) * requestParams.Pagination.PageSize
	}

	return params, filterParams
}

func NewHttpHandler(authSvc iface.IAuthService, tagSvc iface.ITagService) ITagHandler {
	return &handler{
		authSvc: authSvc,
		tagSvc:  tagSvc,
	}
}
