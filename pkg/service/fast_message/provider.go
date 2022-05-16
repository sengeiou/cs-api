package fast_message

import (
	iface "cs-api/pkg/interface"
)

type service struct {
	repo iface.IRepository
}

func NewFastMessageService(Repo iface.IRepository) iface.IFastMessageService {
	return &service{
		repo: Repo,
	}
}
