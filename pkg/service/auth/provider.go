package auth

import (
	"cs-api/config"
	iface "cs-api/pkg/interface"
	_ "embed"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"go.uber.org/fx"
)

type service struct {
	redis  ifaceTool.IRedis
	lua    iface.ILusScript
	repo   iface.IRepository
	config *config.Config
}

type Params struct {
	fx.In

	Cache  ifaceTool.IRedis
	Lua    iface.ILusScript
	Repo   iface.IRepository
	Config *config.Config
}

func NewAuthService(p Params) iface.IAuthService {
	return &service{
		redis:  p.Cache,
		lua:    p.Lua,
		repo:   p.Repo,
		config: p.Config,
	}
}
