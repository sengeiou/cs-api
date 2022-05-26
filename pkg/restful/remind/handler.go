package remind

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
	authSvc   iface.IAuthService
	remindSvc iface.IRemindService
}

type CreateRemindParams struct {
	Title   string       `json:"title" binding:"required"`
	Content string       `json:"content" binding:"required"`
	Status  types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateRemind(c *gin.Context) {
	var (
		err           error
		requestParams CreateRemindParams
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
	params := model.CreateRemindParams{
		Title:     requestParams.Title,
		Content:   requestParams.Content,
		Status:    requestParams.Status,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.remindSvc.CreateRemind(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateRemindParams struct {
	Title   string       `json:"title" binding:"required"`
	Content string       `json:"content" binding:"required"`
	Status  types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateRemind(c *gin.Context) {
	var (
		err           error
		requestParams UpdateRemindParams
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

	params := model.UpdateRemindParams{
		Title:     requestParams.Title,
		Content:   requestParams.Content,
		Status:    requestParams.Status,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        id,
	}

	if err = h.remindSvc.UpdateRemind(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteRemind(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.remindSvc.DeleteRemind(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) GetRemind(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	remind, err := h.remindSvc.GetRemind(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, remind)
}

type ListRemindParams struct {
	Content string       `form:"content" binding:""`
	Status  types.Status `form:"status" binding:"min=0,max=2"`
	types.Pagination
}

func (h *handler) ListRemind(c *gin.Context) {
	var (
		err           error
		requestParams ListRemindParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		fmt.Println(err)
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params, filterParams := formatListRemindParams(requestParams)

	reminds, count, err := h.remindSvc.ListRemind(ctx, params, filterParams)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, reminds, requestParams.Pagination)
}

func formatListRemindParams(requestParams ListRemindParams) (model.ListRemindParams, types.FilterRemindParams) {
	params := model.ListRemindParams{
		Limit:  requestParams.PageSize,
		Offset: 0,
	}

	filterParams := types.FilterRemindParams{
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

func NewHandler(authSvc iface.IAuthService, remindSvc iface.IRemindService) *handler {
	return &handler{
		authSvc:   authSvc,
		remindSvc: remindSvc,
	}
}
