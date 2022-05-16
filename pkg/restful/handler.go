package restful

import (
	"context"
	"crypto/md5"
	"cs-api/pkg"
	"cs-api/pkg/graph/converter"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/go-util/errors"
	ginTool "github.com/golang/go-util/gin"
	"time"
)

func (h *Handler) ListRoomMessage(c *gin.Context) {
	var err error

	clientInfo, err := GetClientInfo(c.Request.Context())
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	messages, err := h.msgSvc.ListRoomMessage(c.Request.Context(), clientInfo.RoomID, "member")
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	tmpResp := make([]*converter.Message, 0)

	for _, message := range messages {
		tmp := converter.Message{
			ID:          message.ID,
			MessageType: converter.MessageTypeDtoMapping[message.Type],
			RoomID:      message.RoomID,
			SenderName:  message.SenderName,
			ContentType: converter.MessageContentTypeDtoMapping[message.ContentType],
			Content:     message.Content,
			Timestamp:   message.Timestamp,
		}

		if message.ExtraInfo != nil {
			tmp.ExtraInfo = &converter.MessageExtraInfo{
				ClientName: message.ExtraInfo.ClientName,
			}
		}

		tmpResp = append(tmpResp, &tmp)
	}

	ginTool.SuccessWithData(c, tmpResp)
}

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
		Type:    pkg.WsClientTypeMember,
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

type UpdateRoomScoreParams struct {
	Score int32 `json:"score" binding:"required,gte=1,lte=5"`
}

func (h *Handler) UpdateRoomScore(c *gin.Context) {
	var err error

	clientInfo, err := GetClientInfo(c.Request.Context())
	if err != nil {
		ginTool.ErrorAuth(c)
		return
	}

	var params UpdateRoomScoreParams
	if err = c.ShouldBindJSON(&params); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	if err = h.roomSvc.UpdateRoomScore(c.Request.Context(), clientInfo.RoomID, params.Score); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

// UploadFile 上傳檔案
func (h *Handler) UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	if file == nil {
		ginTool.Error(c, fmt.Errorf("no file uploaded"))
		return
	}

	src, err := file.Open()
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	defer src.Close()

	url, err := h.storage.Upload(c.Request.Context(), src, file.Filename)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, gin.H{"url": url})
}

func GetClientInfo(ctx context.Context) (pkg.ClientInfo, error) {
	clientInfo := ctx.Value("client_info").(pkg.ClientInfo)
	if clientInfo.ID == 0 {
		return pkg.ClientInfo{}, errors.ErrorAuth
	}

	return clientInfo, nil
}
