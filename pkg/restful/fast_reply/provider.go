package fast_reply

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewService,
		NewHandler,
	),
	fx.Invoke(
		InitTransport,
	),
)
