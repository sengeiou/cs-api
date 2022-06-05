package member

import (
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
)

type handler struct {
	authSvc   iface.IAuthService
	memberSvc iface.IMemberService
}

type ListMemberParams struct {
	Mobile string `form:"mobile" binding:""`
	Email  string `form:"email" binding:""`
	types.Pagination
}

func (h *handler) ListMember(c *gin.Context) {
	var (
		err           error
		requestParams ListMemberParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params, filterParams := formatListMemberParams(requestParams)

	reminds, count, err := h.memberSvc.ListMember(ctx, params, filterParams)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, reminds, requestParams.Pagination)
}

func formatListMemberParams(requestParams ListMemberParams) (model.ListMemberParams, types.FilterMemberParams) {
	params := model.ListMemberParams{
		Limit:  requestParams.PageSize,
		Offset: 0,
	}

	filterParams := types.FilterMemberParams{
		Mobile: nil,
		Email:  nil,
	}

	if requestParams.Mobile != "" {
		filterParams.Mobile = &requestParams.Mobile
	}

	if requestParams.Email != "" {
		filterParams.Email = &requestParams.Email
	}

	if requestParams.Page > 0 {
		params.Offset = (requestParams.Page - 1) * requestParams.PageSize
	}

	return params, filterParams
}

func (h *handler) GetOnlineStatus(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	status, err := h.memberSvc.GetOnlineStatus(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, gin.H{
		"status": status,
	})
}

func NewHandler(authSvc iface.IAuthService, memberSvc iface.IMemberService) *handler {
	return &handler{
		authSvc:   authSvc,
		memberSvc: memberSvc,
	}
}
