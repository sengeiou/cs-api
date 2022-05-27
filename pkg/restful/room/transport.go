package room

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

	routes.GET("/rooms",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.AuthSvc.CheckPermission("ListRoom"),
		p.R.Op("ListRoom"),
		p.H.ListRoom,
	)

	routes.GET("/staff-rooms",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.R.Op("ListStaffRoom"),
		p.H.ListStaffRoom,
	)

	routes.PATCH("/room/accept/:id",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.R.Op("AcceptRoom"),
		p.H.AcceptRoom,
	)

	routes.PATCH("/room/close/:id",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.R.Op("CloseRoom"),
		p.H.CloseRoom,
	)

	routes.PATCH("/room/transfer/:id",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeStaff),
		p.AuthSvc.CheckPermission("TransferRoom"),
		p.R.Op("TransferRoom"),
		p.H.TransferRoom,
	)

	routes.PATCH("/room/score",
		p.AuthSvc.SetClientInfo(pkg.ClientTypeMember),
		p.R.Op("UpdateRoomScore"),
		p.H.UpdateRoomScore,
	)

	// TODO: May check ip white list for this api
	routes.POST("/room", p.H.CreateRoom)
}
