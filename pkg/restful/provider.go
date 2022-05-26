package restful

import (
	iface "cs-api/pkg/interface"
	iface2 "github.com/AndySu1021/go-util/interface"
	"github.com/gin-gonic/gin"
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

//var Module = fx.Options(
//	fx.Provide(
//		NewHandler,
//	),
//	fx.Invoke(
//		InitHandler,
//	),
//)

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
	engine.POST("/api/member/room", h.CreateRoom)
}
