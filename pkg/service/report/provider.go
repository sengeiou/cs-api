package report

import (
	iface "cs-api/pkg/interface"
)

type service struct {
	repo iface.IRepository
}

func NewReportService(Repo iface.IRepository) iface.IReportService {
	return &service{
		repo: Repo,
	}
}
