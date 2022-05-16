package notice

import (
	iface "cs-api/pkg/interface"
)

type service struct {
	repo iface.IRepository
}

func NewNoticeService(Repo iface.IRepository) iface.INoticeService {
	return &service{
		repo: Repo,
	}
}
