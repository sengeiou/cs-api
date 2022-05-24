package ws

import (
	"cs-api/pkg"
	"encoding/json"
	ginTool "github.com/AndySu1021/go-util/gin"
	iface "github.com/AndySu1021/go-util/interface"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type RequestParams struct {
	Type pkg.ClientType `form:"type" binding:"required"`
	SID  string         `form:"sid" binding:"required"`
}

func (h *Handler) ChatHandler(c *gin.Context) {
	var req RequestParams
	if err := c.BindQuery(&req); err != nil {
		ginTool.Error(c, err)
		return
	}

	var redisKey string
	if req.Type == pkg.ClientTypeStaff {
		redisKey = "token:staff:" + req.SID
	} else if req.Type == pkg.ClientTypeMember {
		redisKey = "token:member:" + req.SID
	}

	payload, err := h.redis.Get(c.Request.Context(), redisKey)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	var clientInfo pkg.ClientInfo
	if err = json.Unmarshal([]byte(payload), &clientInfo); err != nil {
		ginTool.Error(c, err)
		return
	}

	// upgrade http protocol to websocket protocol
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://127.0.0.1:8080" || origin == "http://localhost:9528"
	}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		_ = conn.Close()
		return
	}

	clientInfo.Conn = conn

	h.manager.Register(clientInfo)
}

type Handler struct {
	redis   iface.IRedis
	manager *ClientManager
}

func NewHandler(redis iface.IRedis, manager *ClientManager) *Handler {
	return &Handler{
		redis:   redis,
		manager: manager,
	}
}
