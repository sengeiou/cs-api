package member

import (
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/restful/tool"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Engine  *gin.Engine
	AuthSvc iface.IAuthService
	H       *handler
	R       *tool.RequestInstrument
}

func InitTransport(p Params) {
	routes := p.Engine.Group("/api", p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff))

	routes.GET("/members",
		p.AuthSvc.CheckPermission("ListMember"),
		p.R.Op("ListMember"),
		p.H.ListMember,
	)

	routes.GET("/member/:id/online-status",
		p.R.Op("GetOnlineStatus"),
		p.H.GetOnlineStatus,
	)
}
