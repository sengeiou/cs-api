package service

import (
	"cs-api/pkg/service/room"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		room.NewRoomService,
	),
)
