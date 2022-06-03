package chat

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewNotifier,
		NewStaffDispatcher,
		NewClientManager,
		NewHandler,
		NewRedisWorker,
	),
	fx.Invoke(
		InitClientManager,
		InitTransport,
		InitRedisSubscriber,
	),
)
