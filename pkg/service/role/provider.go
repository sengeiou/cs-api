package role

import (
	iface "cs-api/pkg/interface"
)

type service struct {
	repo iface.IRepository
}

func NewRoleService(Repo iface.IRepository) iface.IRoleService {
	return &service{
		repo: Repo,
	}
}
