package room

import (
	iface "cs-api/pkg/interface"
	ifaceTool "github.com/golang/go-util/interface"
	"go.uber.org/fx"
)

type service struct {
	redis     ifaceTool.IRedis
	lua       iface.ILusScript
	memberSvc iface.IMemberService
	repo      iface.IRepository
}

type Params struct {
	fx.In

	Redis     ifaceTool.IRedis
	Lua       iface.ILusScript
	MemberSvc iface.IMemberService
	Repo      iface.IRepository
}

func NewRoomService(p Params) iface.IRoomService {
	return &service{
		redis:     p.Redis,
		lua:       p.Lua,
		memberSvc: p.MemberSvc,
		repo:      p.Repo,
	}
}
