package lua

import (
	iface "cs-api/pkg/interface"
	_ "embed"
	"github.com/go-redis/redis/v8"
	iface2 "github.com/AndySu1021/go-util/interface"
)

//go:embed removeToken.lua
var removeTokenScript string

//go:embed setToken.lua
var setTokenScript string

type script struct {
	redisClient iface2.IRedis
	removeToken *redis.Script
	setToken    *redis.Script
}

func NewLua(redisClient iface2.IRedis) iface.ILusScript {
	return &script{
		redisClient: redisClient,
		removeToken: redis.NewScript(removeTokenScript),
		setToken:    redis.NewScript(setTokenScript),
	}
}
