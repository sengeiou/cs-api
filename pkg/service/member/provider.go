package member

import (
	iface "cs-api/pkg/interface"
)

type service struct {
	repo iface.IRepository
}

func NewMemberService(Repo iface.IRepository) iface.IMemberService {
	return &service{
		repo: Repo,
	}
}
