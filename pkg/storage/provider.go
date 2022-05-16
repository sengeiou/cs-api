package storage

import (
	"cs-api/config"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"errors"
)

func NewStorage(config *config.DiskConfig) (iface.IStorage, error) {
	switch config.Driver {
	case types.DiskDriverLocal:
		return &DiskLocal{RootDir: "./tmp/", BaseUrl: config.BaseUrl}, nil
	case types.DiskDriverS3:
		return &DiskS3{}, nil
	}
	return nil, errors.New("driver not support")
}
