package restful

import (
	iface "cs-api/pkg/interface"
	"cs-api/pkg/restful/auth"
	"cs-api/pkg/restful/cs_config"
	"cs-api/pkg/restful/fast_reply"
	"cs-api/pkg/restful/member"
	"cs-api/pkg/restful/message"
	"cs-api/pkg/restful/notice"
	"cs-api/pkg/restful/remind"
	"cs-api/pkg/restful/report"
	"cs-api/pkg/restful/role"
	"cs-api/pkg/restful/staff"
	"cs-api/pkg/restful/tag"
	"cs-api/pkg/restful/tool"
	iface2 "github.com/AndySu1021/go-util/interface"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	fx.Provide(
		tool.NewRequestInstrument,
		NewH,
	),
	fx.Invoke(
		InitCommonHandler,
		InitH,
	),
	auth.Module,
	tag.Module,
	role.Module,
	notice.Module,
	remind.Module,
	member.Module,
	cs_config.Module,
	report.Module,
	staff.Module,
	message.Module,
	fast_reply.Module,
)

func InitCommonHandler(engine *gin.Engine) {
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	engine.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })
}

type Handler struct {
	authSvc iface.IAuthService
	roomSvc iface.IRoomService
	msgSvc  iface.IMessageService
	redis   iface2.IRedis
	lua     iface.ILusScript
	storage iface.IStorage
}

type Params struct {
	fx.In

	AuthSvc iface.IAuthService
	RoomSvc iface.IRoomService
	MsgSvc  iface.IMessageService
	Redis   iface2.IRedis
	Lua     iface.ILusScript
	Storage iface.IStorage
}

func NewH(p Params) *Handler {
	return &Handler{
		authSvc: p.AuthSvc,
		roomSvc: p.RoomSvc,
		msgSvc:  p.MsgSvc,
		redis:   p.Redis,
		lua:     p.Lua,
		storage: p.Storage,
	}
}

func InitH(engine *gin.Engine, h *Handler) {
	engine.POST("/api/member/room", h.CreateRoom)
}
