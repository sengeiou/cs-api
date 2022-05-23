package message

import (
	iface2 "cs-api/pkg/interface"
	iface "github.com/AndySu1021/go-util/interface"
)

type service struct {
	repo iface.IMongoRepository
}

func NewMessageService(Repo iface.IMongoRepository) iface2.IMessageService {
	return &service{
		repo: Repo,
	}
}
