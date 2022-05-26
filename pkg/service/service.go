package service

import (
	"cs-api/pkg/service/fast_message"
	"cs-api/pkg/service/message"
	"cs-api/pkg/service/report"
	"cs-api/pkg/service/room"
	"cs-api/pkg/service/staff"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		staff.NewStaffService,
		room.NewRoomService,
		fast_message.NewFastMessageService,
		report.NewReportService,
		message.NewMessageService,
	),
)
