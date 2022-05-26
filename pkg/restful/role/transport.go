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

	routes.GET("/roles", p.AuthSvc.CheckPermission("Role.List"), p.R.Op("ListRole"), p.H.ListRole)
	routes.POST("/role", p.AuthSvc.CheckPermission("Role.Create"), p.R.Op("CreateRole"), p.H.CreateRole)
	routes.GET("/role/:id", p.AuthSvc.CheckPermission("Role.Get"), p.R.Op("GetRole"), p.H.GetRole)
	routes.PUT("/role/:id", p.AuthSvc.CheckPermission("Role.Update"), p.R.Op("UpdateRole"), p.H.UpdateRole)
	routes.DELETE("/role/:id", p.AuthSvc.CheckPermission("Role.Delete"), p.R.Op("DeleteRole"), p.H.DeleteRole)
}
