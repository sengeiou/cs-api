package fast_reply

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

	routes.GET("/fast-replies",
		p.AuthSvc.CheckPermission("ListFastReply"),
		p.R.Op("ListFastReply"),
		p.H.ListFastReply,
	)

	routes.POST("/fast-reply",
		p.AuthSvc.CheckPermission("CreateFastReply"),
		p.R.Op("CreateFastReply"),
		p.H.CreateFastReply,
	)

	routes.GET("/fast-reply/:id",
		p.AuthSvc.CheckPermission("GetFastReply"),
		p.R.Op("GetFastReply"),
		p.H.GetFastReply,
	)

	routes.PUT("/fast-reply/:id",
		p.AuthSvc.CheckPermission("UpdateFastReply"),
		p.R.Op("UpdateFastReply"),
		p.H.UpdateFastReply,
	)

	routes.DELETE("/fast-reply/:id",
		p.AuthSvc.CheckPermission("DeleteFastReply"),
		p.R.Op("DeleteFastReply"),
		p.H.DeleteFastReply,
	)

	routes.GET("/fast-reply/group",
		p.AuthSvc.CheckPermission("ListFastReplyGroup"),
		p.R.Op("ListFastReplyGroup"),
		p.H.ListFastReplyGroup,
	)

	routes.GET("/fast-reply/categories",
		p.AuthSvc.CheckPermission("ListFastReplyCategory"),
		p.R.Op("ListFastReplyCategory"),
		p.H.ListFastReplyCategory,
	)

	routes.POST("/fast-reply/category",
		p.AuthSvc.CheckPermission("CreateFastReplyCategory"),
		p.R.Op("CreateFastReplyCategory"),
		p.H.CreateFastReplyCategory,
	)
}
