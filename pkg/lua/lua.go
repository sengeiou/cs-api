package lua

import (
	"context"
	"time"
)

func (s *script) RemoveToken(ctx context.Context, clientType string, name string) error {
	return s.removeToken.Run(ctx, s.redisClient.GetClient(), []string{clientType, name}).Err()
}

func (s *script) SetToken(ctx context.Context, clientType string, name string, token string, value interface{}, duration time.Duration) error {
	expire := int64(duration / time.Second)
	return s.setToken.Run(ctx, s.redisClient.GetClient(), []string{clientType, name, token}, value, expire).Err()
}
