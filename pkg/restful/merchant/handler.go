package merchant

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
	authSvc     iface.IAuthService
	merchantSvc iface.IMerchantService
}

type CreateMerchantParams struct {
	Name   string       `json:"name" binding:"required"`
	Code   string       `json:"code" binding:"required"`
	Status types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) CreateMerchant(c *gin.Context) {
	var (
		err           error
		requestParams CreateMerchantParams
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
	params := model.CreateMerchantParams{
		Name:      requestParams.Name,
		Code:      requestParams.Code,
		Status:    requestParams.Status,
		CreatedBy: staffInfo.ID,
		CreatedAt: now,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: now,
	}

	if err = h.merchantSvc.CreateMerchant(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateMerchantParams struct {
	Name   string       `json:"name" binding:"required"`
	Code   string       `json:"code" binding:"required"`
	Status types.Status `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateMerchant(c *gin.Context) {
	var (
		err           error
		requestParams UpdateMerchantParams
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

	params := model.UpdateMerchantParams{
		Name:      requestParams.Name,
		Code:      requestParams.Code,
		Status:    requestParams.Status,
		UpdatedBy: staffInfo.ID,
		UpdatedAt: time.Now().UTC(),
		ID:        id,
	}

	if err = h.merchantSvc.UpdateMerchant(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) DeleteMerchant(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.merchantSvc.DeleteMerchant(ctx, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func (h *handler) ListAvailableMerchant(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)

	tags, err := h.merchantSvc.ListAvailableMerchant(ctx)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, tags)
}

func (h *handler) GetMerchant(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	merchant, err := h.merchantSvc.GetMerchant(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, merchant)
}

type ListMerchantParams struct {
	Name   string       `form:"name" binding:""`
	Status types.Status `form:"status" binding:"min=0,max=2"`
	types.Pagination
}

func (h *handler) ListMerchant(c *gin.Context) {
	var (
		err           error
		requestParams ListMerchantParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params, filterParams := formatListMerchantParams(requestParams)

	merchants, count, err := h.merchantSvc.ListMerchant(ctx, params, filterParams)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, merchants, requestParams.Pagination)
}

func formatListMerchantParams(requestParams ListMerchantParams) (model.ListMerchantParams, types.FilterMerchantParams) {
	params := model.ListMerchantParams{
		Limit:  requestParams.Pagination.PageSize,
		Offset: 0,
	}

	filterParams := types.FilterMerchantParams{
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

func NewHandler(authSvc iface.IAuthService, merchantSvc iface.IMerchantService) *handler {
	return &handler{
		authSvc:     authSvc,
		merchantSvc: merchantSvc,
	}
}
