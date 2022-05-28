package room

import (
	"crypto/md5"
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"encoding/json"
	"fmt"
	errorTool "github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type handler struct {
	authSvc iface.IAuthService
	roomSvc iface.IRoomService
	lua     iface.ILusScript
}

type ListRoomParams struct {
	RoomID  int64            `form:"room_id" binding:""`
	StaffID int64            `form:"staff_id" binding:""`
	Status  types.RoomStatus `form:"status" binding:"min=0,max=3"`
	types.Pagination
}

func (h *handler) ListRoom(c *gin.Context) {
	var (
		err           error
		requestParams ListRoomParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	params, filterParams := formatListRoomParams(requestParams)

	rooms, count, err := h.roomSvc.ListRoom(ctx, params, filterParams)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, rooms, requestParams.Pagination)
}

func formatListRoomParams(requestParams ListRoomParams) (model.ListRoomParams, types.FilterRoomParams) {
	params := model.ListRoomParams{
		Limit:  requestParams.PageSize,
		Offset: 0,
	}

	filterParams := types.FilterRoomParams{
		RoomID:  nil,
		StaffID: nil,
		Status:  nil,
	}

	if requestParams.RoomID != 0 {
		filterParams.RoomID = &requestParams.RoomID
	}

	if requestParams.StaffID != 0 {
		filterParams.StaffID = &requestParams.StaffID
	}

	if requestParams.Status != types.RoomStatusAll {
		filterParams.Status = &requestParams.Status
	}

	if requestParams.Page > 0 {
		params.Offset = (requestParams.Page - 1) * requestParams.PageSize
	}

	return params, filterParams
}

type ListStaffRoomParams struct {
	Status types.RoomStatus `form:"status" binding:"required,oneof=1 2"`
	types.Pagination
}

func (h *handler) ListStaffRoom(c *gin.Context) {
	var (
		err           error
		requestParams ListStaffRoomParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	params, filterParams := formatListStaffRoomParams(requestParams, staffInfo.ID)

	rooms, count, err := h.roomSvc.ListStaffRoom(ctx, params, filterParams)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	requestParams.Pagination.Total = count

	ginTool.SuccessWithPagination(c, rooms, requestParams.Pagination)
}

func formatListStaffRoomParams(requestParams ListStaffRoomParams, staffId int64) (model.ListStaffRoomParams, types.FilterStaffRoomParams) {
	params := model.ListStaffRoomParams{
		Status: requestParams.Status,
		Limit:  requestParams.PageSize,
		Offset: 0,
	}

	filterParams := types.FilterStaffRoomParams{
		StaffID: nil,
	}

	if requestParams.Status == types.RoomStatusServing {
		filterParams.StaffID = &staffId
	}

	if requestParams.Page > 0 {
		params.Offset = (requestParams.Page - 1) * requestParams.PageSize
	}

	return params, filterParams
}

func (h *handler) AcceptRoom(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
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

	if err = h.roomSvc.AcceptRoom(ctx, staffInfo.ID, id); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type CloseRoomParams struct {
	TagID int64 `json:"tag_id" binding:"required"`
}

func (h *handler) CloseRoom(c *gin.Context) {
	var (
		err           error
		requestParams CloseRoomParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.roomSvc.CloseRoom(ctx, staffInfo.ID, id, requestParams.TagID); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type TransferRoomParams struct {
	StaffID int64 `json:"staff_id" binding:"required"`
}

func (h *handler) TransferRoom(c *gin.Context) {
	var (
		err           error
		requestParams TransferRoomParams
		ctx           = c.Request.Context()
	)

	staffInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeStaff)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.roomSvc.TransferRoom(ctx, staffInfo.ID, id, requestParams.StaffID); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type UpdateRoomScoreParams struct {
	Score int32 `json:"score" binding:"required,min=1,max=5"`
}

func (h *handler) UpdateRoomScore(c *gin.Context) {
	var (
		err           error
		requestParams UpdateRoomScoreParams
		ctx           = c.Request.Context()
	)

	memberInfo, err := h.authSvc.GetClientInfo(ctx, pkg.ClientTypeMember)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	if err = h.roomSvc.UpdateRoomScore(ctx, memberInfo.RoomID, requestParams.Score); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

type CreateRoomParams struct {
	Name     string `json:"name" binding:""`
	DeviceID string `json:"device_id" binding:"required"`
}

func (h *handler) CreateRoom(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	var params CreateRoomParams
	if err = c.ShouldBindJSON(&params); err != nil {
		ginTool.Error(c, errorTool.ErrorValidation)
		return
	}

	room, member, err := h.roomSvc.CreateRoom(ctx, params.DeviceID, params.Name)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = h.lua.RemoveToken(ctx, "member", member.Name); err != nil {
		ginTool.ErrorAuth(c)
		return
	}

	token := fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()+member.Name)))

	clientInfo := pkg.ClientInfo{
		ID:      member.ID,
		Type:    pkg.ClientTypeMember,
		Name:    member.Name,
		RoomID:  room.ID,
		StaffID: room.StaffID,
	}

	payload, _ := json.Marshal(clientInfo)

	if err = h.lua.SetToken(ctx, "member", member.Name, token, payload, 2*time.Hour); err != nil {
		ginTool.ErrorAuth(c)
		return
	}

	ginTool.SuccessWithData(c, fmt.Sprintf("http://127.0.0.1:8080/chat?room_id=%d&name=%s&sid=%s", room.ID, member.Name, token))
}

func NewHandler(authSvc iface.IAuthService, roomSvc iface.IRoomService, lua iface.ILusScript) *handler {
	return &handler{
		authSvc: authSvc,
		roomSvc: roomSvc,
		lua:     lua,
	}
}
