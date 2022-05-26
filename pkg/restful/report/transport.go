package report

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

	routes.GET("/report/daily-tag",
		p.AuthSvc.CheckPermission("DailyTagReport"),
		p.R.Op("DailyTagReport"),
		p.H.DailyTagReport,
	)

	routes.GET("/report/daily-guest",
		p.AuthSvc.CheckPermission("DailyGuestReport"),
		p.R.Op("DailyGuestReport"),
		p.H.DailyGuestReport,
	)
}
