package ws

import (
	"cs-api/pkg"
	"encoding/json"
	"github.com/gin-gonic/gin"
	ginTool "github.com/golang/go-util/gin"
	iface "github.com/golang/go-util/interface"
	"github.com/gorilla/websocket"
	"net/http"
)

type RequestParams struct {
	Type pkg.WsClientType `form:"type" binding:"required"`
	SID  string           `form:"sid" binding:"required"`
}

func (h *Handler) ChatHandler(c *gin.Context) {
	var req RequestParams
	if err := c.BindQuery(&req); err != nil {
		ginTool.Error(c, err)
		return
	}

	var redisKey string
	if req.Type == pkg.WsClientTypeStaff {
		redisKey = "token:staff:" + req.SID
	} else if req.Type == pkg.WsClientTypeMember {
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

//var Count int64
//
//func (h *Handler) TestChatHandler(c *gin.Context) {
//	var req RequestParams
//	if err := c.BindQuery(&req); err != nil {
//		ginTool.Error(c, err)
//		return
//	}
//
//	var redisKey string
//	if req.Type == pkg.WsClientTypeStaff {
//		redisKey = "token:staff:" + req.SID
//	} else if req.Type == pkg.WsClientTypeMember {
//		redisKey = "token:member:" + req.SID
//	}
//	_ = redisKey
//
//	//payload, err := h.redis.Get(c.Request.Context(), redisKey)
//	//if err != nil {
//	//	ginTool.Error(c, err)
//	//	return
//	//}
//
//	atomic.AddInt64(&Count, 1)
//	payload := fmt.Sprintf(`{"id":%d,"type":2,"name":"b","room_id":%d,"staff_id":1,"serving_status":0}`, Count, Count)
//	fmt.Println("Count number is: ", Count)
//
//	var clientInfo pkg.ClientInfo
//	if err := json.Unmarshal([]byte(payload), &clientInfo); err != nil {
//		ginTool.Error(c, err)
//		return
//	}
//
//	// upgrade http protocol to websocket protocol
//	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
//		return true
//	}}).Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		_ = conn.Close()
//		return
//	}
//
//	clientInfo.Conn = conn
//
//	h.manager.Register(clientInfo)
//}

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
