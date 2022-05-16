package service

import (
	"cs-api/pkg/service/auth"
	"cs-api/pkg/service/cs_config"
	"cs-api/pkg/service/fast_message"
	"cs-api/pkg/service/member"
	"cs-api/pkg/service/message"
	"cs-api/pkg/service/notice"
	"cs-api/pkg/service/remind"
	"cs-api/pkg/service/report"
	"cs-api/pkg/service/role"
	"cs-api/pkg/service/room"
	"cs-api/pkg/service/staff"
	"cs-api/pkg/service/tag"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		tag.NewTagService,
		role.NewRoleService,
		staff.NewStaffService,
		room.NewRoomService,
		member.NewMemberService,
		cs_config.NewCsConfigService,
		fast_message.NewFastMessageService,
		report.NewReportService,
		message.NewMessageService,
		auth.NewAuthService,
		notice.NewNoticeService,
		remind.NewRemindService,
	),
)
