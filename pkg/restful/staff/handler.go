package staff

import (
	"cs-api/config"
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/AndySu1021/go-util/helper"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type handler struct {
	authSvc  iface.IAuthService
	staffSvc iface.IStaffService
	config   *config.Config
}

type CreateStaffParams struct {
	RoleID   int64        `json:"role_id" binding:"required"`
	Name     string       `json:"name" binding:"required"`
	Username string       `json:"username" binding:"required"`
	Password string       `json:"password" binding:"required"`
	Status   types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateStaff(c *gin.Context) {
	var (
		err           error
		requestParams CreateStaffParams
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
	params := model.CreateStaffParams{
		RoleID:    requestParams.RoleID,
		Name:      requestParams.Name,
		Username:  requestParams.Username,
		Password:  helper.EncryptPassword(requestParams.Password, h.config.Salt),
		Status:    requestParams.Status,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.staffSvc.CreateStaff(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateStaffParams struct {
	RoleID   int64        `json:"role_id" binding:"required"`
	Name     string       `json:"name" binding:"required"`
	Password string       `json:"password" binding:""`
	Status   types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateStaff(c *gin.Context) {
	var (
		err           error
		requestParams UpdateStaffParams
		params        interface{}
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

	if requestParams.Password != "" {
		params = model.UpdateStaffWithPasswordParams{
			RoleID:    requestParams.RoleID,
			Name:      requestParams.Name,
			Password:  helper.EncryptPassword(requestParams.Password, h.config.Salt),
			Status:    requestParams.Status,
			ID:        id,
			UpdatedBy: staffInfo.ID,
			UpdatedAt: time.Now().UTC(),
		}
	} else {
		params = model.UpdateStaffParams{
			RoleID:    requestParams.RoleID,
			Name:      requestParams.Name,
			Status:    requestParams.Status,
			ID:        id,
			UpdatedBy: staffInfo.ID,
			UpdatedAt: time.Now().UTC(),
		}
	}

	if err = h.staffSvc.UpdateStaff(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteStaff(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.staffSvc.DeleteStaff(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) GetStaff(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	staff, err := h.staffSvc.GetStaff(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, staff)
}

type ListStaffParams struct {
	Name          string                   `form:"name" binding:""`
	Status        types.Status             `form:"status" binding:"min=0,max=2"`
	ServingStatus types.StaffServingStatus `form:"serving_status" binding:"min=0,max=3"`
	types.Pagination
}

func (h *handler) ListStaff(c *gin.Context) {
	var (
		err           error
		requestParams ListStaffParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params, filterParams := formatListStaffParams(requestParams)

	staffs, count, err := h.staffSvc.ListStaff(ctx, params, filterParams)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, staffs, requestParams.Pagination)
}

func formatListStaffParams(requestParams ListStaffParams) (model.ListStaffParams, types.FilterStaffParams) {
	params := model.ListStaffParams{
		Limit:  requestParams.Pagination.PageSize,
		Offset: 0,
	}

	filterParams := types.FilterStaffParams{
		Name:          nil,
		Status:        nil,
		ServingStatus: nil,
	}

	if requestParams.Name != "" {
		filterParams.Name = &requestParams.Name
	}

	if requestParams.Status != types.StatusAll {
		filterParams.Status = &requestParams.Status
	}

	if requestParams.ServingStatus != types.StaffServingStatusAll {
		filterParams.ServingStatus = &requestParams.ServingStatus
	}

	if requestParams.Pagination.Page > 0 {
		params.Offset = (requestParams.Pagination.Page - 1) * requestParams.Pagination.PageSize
	}

	return params, filterParams
}

type UpdateStaffServingStatusParams struct {
	ServingStatus types.StaffServingStatus `json:"serving_status" binding:"required,min=1,max=3"`
}

func (h *handler) UpdateStaffServingStatus(c *gin.Context) {
	var (
		err           error
		requestParams UpdateStaffServingStatusParams
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

	if err = h.staffSvc.UpdateStaffServingStatus(ctx, staffInfo, requestParams.ServingStatus); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateStaffAvatarParams struct {
	Avatar string `json:"avatar" binding:"required"`
}

func (h *handler) UpdateStaffAvatar(c *gin.Context) {
	var (
		err           error
		requestParams UpdateStaffAvatarParams
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

	params := model.UpdateStaffAvatarParams{
		Avatar:    requestParams.Avatar,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        staffInfo.ID,
	}

	if err = h.staffSvc.UpdateStaff(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) ListAvailableStaff(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	staffs, err := h.staffSvc.ListAvailableStaff(ctx, staffInfo.ID)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, staffs)
}

func NewHandler(authSvc iface.IAuthService, staffSvc iface.IStaffService, config *config.Config) *handler {
	return &handler{
		authSvc:  authSvc,
		staffSvc: staffSvc,
		config:   config,
	}
}
