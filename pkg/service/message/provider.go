package message

import (
	iface2 "cs-api/pkg/interface"
	iface "github.com/golang/go-util/interface"
)

type service struct {
	repo iface.IMongoRepository
}

func NewMessageService(Repo iface.IMongoRepository) iface2.IMessageService {
	return &service{
		repo: Repo,
	}
}
