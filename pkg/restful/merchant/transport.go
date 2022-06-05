package merchant

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

	routes.GET("/merchants",
		p.AuthSvc.CheckPermission("ListMerchant"),
		p.R.Op("ListMerchant"),
		p.H.ListMerchant,
	)

	routes.POST("/merchant",
		p.AuthSvc.CheckPermission("CreateMerchant"),
		p.R.Op("CreateMerchant"),
		p.H.CreateMerchant,
	)

	routes.GET("/merchant/:id",
		p.AuthSvc.CheckPermission("GetMerchant"),
		p.R.Op("GetMerchant"),
		p.H.GetMerchant,
	)

	routes.PUT("/merchant/:id",
		p.AuthSvc.CheckPermission("UpdateMerchant"),
		p.R.Op("UpdateMerchant"),
		p.H.UpdateMerchant,
	)

	routes.DELETE("/merchant/:id",
		p.AuthSvc.CheckPermission("DeleteMerchant"),
		p.R.Op("DeleteMerchant"),
		p.H.DeleteMerchant,
	)

	routes.GET("/available-merchants",
		p.R.Op("ListAvailableMerchant"),
		p.H.ListAvailableMerchant,
	)
}
