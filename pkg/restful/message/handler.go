package message

import (
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	authSvc    iface.IAuthService
	messageSvc iface.IMessageService
}

type ListMessageParams struct {
	RoomID  int64  `form:"room_id" binding:""`
	StaffID int64  `form:"staff_id" binding:""`
	Content string `form:"content" binding:""`
	types.Pagination
}

func (h *handler) ListMessage(c *gin.Context) {
	var (
		err           error
		requestParams ListMessageParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params := types.ListMessageParams{
		RoomID:   requestParams.RoomID,
		StaffID:  requestParams.StaffID,
		Content:  requestParams.Content,
		Page:     int64(requestParams.Pagination.Page),
		PageSize: int64(requestParams.Pagination.PageSize),
	}

	messages, count, err := h.messageSvc.ListMessage(ctx, params)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, messages, requestParams.Pagination)
}

type ListStaffRoomMessageParams struct {
	RoomID int64 `form:"room_id" binding:""`
}

func (h *handler) ListStaffRoomMessage(c *gin.Context) {
	var (
		err           error
		requestParams ListStaffRoomMessageParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	messages, err := h.messageSvc.ListRoomMessage(ctx, requestParams.RoomID, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, messages)
}

func (h *handler) ListMemberRoomMessage(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	clientInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeMember)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	messages, err := h.messageSvc.ListRoomMessage(ctx, clientInfo.RoomID, pkg.ClientTypeMember)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, messages)
}

func NewHandler(authSvc iface.IAuthService, messageSvc iface.IMessageService) *handler {
	return &handler{
		authSvc:    authSvc,
		messageSvc: messageSvc,
	}
}
