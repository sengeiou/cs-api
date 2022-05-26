package service

import (
	"cs-api/pkg/service/fast_message"
	"cs-api/pkg/service/message"
	"cs-api/pkg/service/room"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		room.NewRoomService,
		fast_message.NewFastMessageService,
		message.NewMessageService,
	),
)
