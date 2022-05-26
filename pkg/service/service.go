package service

import (
	"cs-api/pkg/service/cs_config"
	"cs-api/pkg/service/fast_message"
	"cs-api/pkg/service/member"
	"cs-api/pkg/service/message"
	"cs-api/pkg/service/notice"
	"cs-api/pkg/service/remind"
	"cs-api/pkg/service/report"
	"cs-api/pkg/service/room"
	"cs-api/pkg/service/staff"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		staff.NewStaffService,
		room.NewRoomService,
		member.NewMemberService,
		cs_config.NewCsConfigService,
		fast_message.NewFastMessageService,
		report.NewReportService,
		message.NewMessageService,
		notice.NewNoticeService,
		remind.NewRemindService,
	),
)
