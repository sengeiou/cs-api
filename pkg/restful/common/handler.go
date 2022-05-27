package common

import (
	ginTool "github.com/AndySu1021/go-util/gin"
	iface "github.com/AndySu1021/go-util/interface"
	"github.com/gin-gonic/gin"
)

type handler struct {
	storage iface.IStorage
}

func (h *handler) UploadFile(c *gin.Context) {
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

func NewHandler(storage iface.IStorage) *handler {
	return &handler{
		storage: storage,
	}
}
