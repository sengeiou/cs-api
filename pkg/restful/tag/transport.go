package tag

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

	routes.GET("/tags", p.AuthSvc.CheckPermission("Tag.List"), p.R.Op("ListTag"), p.H.ListTag)
	routes.POST("/tag", p.AuthSvc.CheckPermission("Tag.Create"), p.R.Op("CreateTag"), p.H.CreateTag)
	routes.GET("/tag/:id", p.AuthSvc.CheckPermission("Tag.Get"), p.R.Op("GetTag"), p.H.GetTag)
	routes.PUT("/tag/:id", p.AuthSvc.CheckPermission("Tag.Update"), p.R.Op("UpdateTag"), p.H.UpdateTag)
	routes.DELETE("/tag/:id", p.AuthSvc.CheckPermission("Tag.Delete"), p.R.Op("DeleteTag"), p.H.DeleteTag)
}
