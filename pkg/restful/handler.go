package restful

import (
	"crypto/md5"
	"cs-api/pkg"
	"encoding/json"
	"fmt"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateRoomParams struct {
	Name     string `json:"name" binding:""`
	DeviceID string `json:"device_id" binding:"required"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var err error

	var params CreateRoomParams
	if err = c.ShouldBindJSON(&params); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	ctx := c.Request.Context()

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
