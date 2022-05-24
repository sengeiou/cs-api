package restful

import (
	"context"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"encoding/json"
	gin2 "github.com/AndySu1021/go-util/gin"
	iface2 "github.com/AndySu1021/go-util/interface"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
)

type Handler struct {
	authSvc iface.IAuthService
	roomSvc iface.IRoomService
	msgSvc  iface.IMessageService
	redis   iface2.IRedis
	lua     iface.ILusScript
	storage iface.IStorage
}

var Module = fx.Options(
	fx.Provide(
		NewHandler,
	),
	fx.Invoke(
		InitHandler,
	),
)

type Params struct {
	fx.In

	AuthSvc iface.IAuthService
	RoomSvc iface.IRoomService
	MsgSvc  iface.IMessageService
	Redis   iface2.IRedis
	Lua     iface.ILusScript
	Storage iface.IStorage
}

func NewHandler(p Params) *Handler {
	return &Handler{
		authSvc: p.AuthSvc,
		roomSvc: p.RoomSvc,
		msgSvc:  p.MsgSvc,
		redis:   p.Redis,
		lua:     p.Lua,
		storage: p.Storage,
	}
}

func InitHandler(engine *gin.Engine, h *Handler) {
	engine.POST("/api/auth/login", h.Login)
	engine.GET("/api/member/message/list", CheckMemberToken(h), h.ListRoomMessage)
	engine.POST("/api/member/room/create", h.CreateRoom)
	engine.POST("/api/member/room/score", CheckMemberToken(h), h.UpdateRoomScore)
	engine.POST("/api/member/upload", CheckMemberToken(h), h.UploadFile)
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

func CheckMemberToken(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")
		if token == "" {
			gin2.ErrorAuth(c)
			c.Abort()
			return
		}

		result, err := h.redis.Get(c.Request.Context(), "token:member:"+token)
		if err != nil {
			gin2.ErrorAuth(c)
			c.Abort()
			return
		}

		var tmp pkg.ClientInfo
		err = json.Unmarshal([]byte(result), &tmp)
		if err != nil {
			gin2.Error(c, err)
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), "client_info", tmp)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
