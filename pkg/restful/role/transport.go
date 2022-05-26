package role

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

	routes.GET("/roles",
		p.AuthSvc.CheckPermission("ListRole"),
		p.R.Op("ListRole"),
		p.H.ListRole,
	)

	routes.POST("/role",
		p.AuthSvc.CheckPermission("CreateRole"),
		p.R.Op("CreateRole"),
		p.H.CreateRole,
	)

	routes.GET("/role/:id",
		p.AuthSvc.CheckPermission("GetRole"),
		p.R.Op("GetRole"),
		p.H.GetRole,
	)

	routes.PUT("/role/:id",
		p.AuthSvc.CheckPermission("UpdateRole"),
		p.R.Op("UpdateRole"),
		p.H.UpdateRole,
	)

	routes.DELETE("/role/:id",
		p.AuthSvc.CheckPermission("DeleteRole"),
		p.R.Op("DeleteRole"),
		p.H.DeleteRole,
	)
}
