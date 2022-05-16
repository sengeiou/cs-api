package iface

import (
	"context"
	"time"
)

type ILusScript interface {
	RemoveToken(ctx context.Context, clientType string, name string) error
	SetToken(ctx context.Context, clientType string, name string, token string, value interface{}, duration time.Duration) error
}
