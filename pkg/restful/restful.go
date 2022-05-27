package restful

import (
	"cs-api/pkg/restful/auth"
	"cs-api/pkg/restful/common"
	"cs-api/pkg/restful/cs_config"
	"cs-api/pkg/restful/fast_reply"
	"cs-api/pkg/restful/member"
	"cs-api/pkg/restful/message"
	"cs-api/pkg/restful/notice"
	"cs-api/pkg/restful/remind"
	"cs-api/pkg/restful/report"
	"cs-api/pkg/restful/role"
	"cs-api/pkg/restful/room"
	"cs-api/pkg/restful/staff"
	"cs-api/pkg/restful/tag"
	"cs-api/pkg/restful/tool"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		tool.NewRequestInstrument,
	),
	auth.Module,
	tag.Module,
	role.Module,
	notice.Module,
	remind.Module,
	member.Module,
	cs_config.Module,
	report.Module,
	staff.Module,
	message.Module,
	fast_reply.Module,
	room.Module,
	common.Module,
)
