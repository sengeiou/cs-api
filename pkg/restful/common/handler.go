package common

import (
	"fmt"
	ginTool "github.com/AndySu1021/go-util/gin"
	iface "github.com/AndySu1021/go-util/interface"
	"github.com/gin-gonic/gin"
)

type handler struct {
	redis   iface.IRedis
	storage iface.IStorage
}

type UploadFileParams struct {
	Type string `form:"type" binding:"required,oneof=member staff"`
}

func (h *handler) UploadFile(c *gin.Context) {
	token := c.GetHeader("X-Token")

	var req UploadFileParams
	if err := c.ShouldBind(&req); err != nil {
		ginTool.Error(c, err)
		return
	}

	redisKey := fmt.Sprintf("token:%s:%s", req.Type, token)
	payload, err := h.redis.Get(c.Request.Context(), redisKey)
	if err != nil || payload == "" {
		ginTool.ErrorAuth(c)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	tmp, err := file.Open()
	if err != nil {
		ginTool.Error(c, err)
		return
	}
	defer tmp.Close()

	url, err := h.storage.Upload(c.Request.Context(), tmp, file.Filename)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, url)
}

func NewHandler(redis iface.IRedis, storage iface.IStorage) *handler {
	return &handler{
		redis:   redis,
		storage: storage,
	}
}
