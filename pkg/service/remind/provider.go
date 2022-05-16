package remind

import (
	iface "cs-api/pkg/interface"
)

type service struct {
	repo iface.IRepository
}

func NewRemindService(Repo iface.IRepository) iface.IRemindService {
	return &service{
		repo: Repo,
	}
}
