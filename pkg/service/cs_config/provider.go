package cs_config

import (
	iface "cs-api/pkg/interface"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	"go.uber.org/fx"
)

type service struct {
	redis ifaceTool.IRedis
	repo  iface.IRepository
}

type Params struct {
	fx.In

	Redis ifaceTool.IRedis
	Repo  iface.IRepository
}

func NewCsConfigService(p Params) iface.ICsConfigService {
	return &service{
		redis: p.Redis,
		repo:  p.Repo,
	}
}
