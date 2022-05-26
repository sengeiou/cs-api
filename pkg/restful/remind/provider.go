package remind

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewService,
		NewHandler,
	),
	fx.Invoke(
		InitTransport,
	),
)
