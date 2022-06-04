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
	routes := p.Engine.Group("/api")

	routes.GET("/member/:id/status",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.R.Op("GetMemberStatus"),
		p.H.GetMemberStatus,
	)
}
