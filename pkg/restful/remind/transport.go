package remind

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

	routes.GET("/reminds",
		p.AuthSvc.CheckPermission("ListRemind"),
		p.R.Op("ListRemind"),
		p.H.ListRemind,
	)

	routes.POST("/remind",
		p.AuthSvc.CheckPermission("CreateRemind"),
		p.R.Op("CreateRemind"),
		p.H.CreateRemind,
	)

	routes.GET("/remind/:id",
		p.AuthSvc.CheckPermission("GetRemind"),
		p.R.Op("GetRemind"),
		p.H.GetRemind,
	)

	routes.PUT("/remind/:id",
		p.AuthSvc.CheckPermission("UpdateRemind"),
		p.R.Op("UpdateRemind"),
		p.H.UpdateRemind,
	)

	routes.DELETE("/remind/:id",
		p.AuthSvc.CheckPermission("DeleteRemind"),
		p.R.Op("DeleteRemind"),
		p.H.DeleteRemind,
	)

	routes.GET("/active-reminds",
		p.R.Op("ListActiveRemind"),
		p.H.ListActiveRemind,
	)
}
