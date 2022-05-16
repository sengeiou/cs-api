package tag

import iface "cs-api/pkg/interface"

type service struct {
	repo iface.IRepository
}

func NewTagService(Repo iface.IRepository) iface.ITagService {
	return &service{
		repo: Repo,
	}
}
