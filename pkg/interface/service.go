package iface

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	model2 "cs-api/pkg/model"
	"cs-api/pkg/types"
	"github.com/gin-gonic/gin"
	"time"
)

type IAuthService interface {
	Login(ctx context.Context, username, password string) (pkg.ClientInfo, error)
	Logout(ctx context.Context, userInfo pkg.ClientInfo) error
	SetClientInfo(clientType pkg.ClientType) gin.HandlerFunc
	GetClientInfo(ctx context.Context, clientType pkg.ClientType) (pkg.ClientInfo, error)
	CheckPermission(permission string) gin.HandlerFunc
}

type IStaffService interface {
	ListStaff(ctx context.Context, params model.ListStaffParams, filterParams types.FilterStaffParams) ([]model.ListStaffRow, int64, error)
	GetStaff(ctx context.Context, staffId int64) (model.GetStaffRow, error)
	CreateStaff(ctx context.Context, params model.CreateStaffParams) error
	UpdateStaff(ctx context.Context, params interface{}) error
	DeleteStaff(ctx context.Context, staffId int64) error
	UpdateStaffServingStatus(ctx context.Context, staffInfo pkg.ClientInfo, status types.StaffServingStatus) error
	ListAvailableStaff(ctx context.Context, staffId int64) ([]model.Staff, error)
}

type IRoomService interface {
	CreateRoom(ctx context.Context, deviceId string, name string) (model.Room, model.Member, error)
	AcceptRoom(ctx context.Context, staffId int64, roomId int64) error
	CloseRoom(ctx context.Context, roomId int64, tagId int64) error
	TransferRoom(ctx context.Context, roomId int64, staffId int64) error
	UpdateRoomScore(ctx context.Context, roomId int64, score int32) error
	ListStaffRoom(ctx context.Context, params model.ListStaffRoomParams, filterParams types.FilterStaffRoomParams) ([]model.ListStaffRoomRow, int64, error)
	ListRoom(ctx context.Context, params model.ListRoomParams, filterParams types.FilterRoomParams) ([]model.ListRoomRow, int64, error)
	GetStaffRooms(ctx context.Context, staffId int64) ([]int64, error)
}

type IMessageService interface {
	CreateMessage(ctx context.Context, message model2.Message) error
	ListRoomMessage(ctx context.Context, roomId int64, clientType pkg.ClientType) ([]model2.Message, error)
	ListMessage(ctx context.Context, params types.ListMessageParams) ([]model2.Message, int64, error)
}

type ITagService interface {
	ListTag(ctx context.Context, params model.ListTagParams, filterParams types.FilterTagParams) ([]model.ListTagRow, int64, error)
	GetTag(ctx context.Context, tagId int64) (model.GetTagRow, error)
	CreateTag(ctx context.Context, params model.CreateTagParams) error
	UpdateTag(ctx context.Context, params model.UpdateTagParams) error
	DeleteTag(ctx context.Context, tagId int64) error
}

type IFastMessageService interface {
	ListFastMessage(ctx context.Context, params model.ListFastMessageParams, filterParams types.FilterFastMessageParams) ([]model.ListFastMessageRow, int64, error)
	GetFastMessage(ctx context.Context, id int64) (model.FastMessage, error)
	CreateFastMessage(ctx context.Context, params model.CreateFastMessageParams) error
	UpdateFastMessage(ctx context.Context, params model.UpdateFastMessageParams) error
	DeleteFastMessage(ctx context.Context, id int64) error
	ListCategory(ctx context.Context) ([]model.Constant, error)
	CreateCategory(ctx context.Context, params model.CreateFastMessageCategoryParams) error
	ListFastMessageGroup(ctx context.Context) ([]pkg.FastMessageGroupItem, error)
}

type IReportService interface {
	DailyTagReport(ctx context.Context, startDate, endDate time.Time) ([]pkg.DailyTagReportColumn, map[string]map[string]int32, error)
	DailyGuestReport(ctx context.Context, startDate, endDate time.Time) (map[string]int32, error)
}

type ICsConfigService interface {
	GetCsConfig(ctx context.Context) (types.CsConfig, error)
	UpdateCsConfig(ctx context.Context, staffId int64, config types.CsConfig) error
}

type IMemberService interface {
	GetOrCreateMember(ctx context.Context, name string, deviceId string) (model.Member, error)
}

type IRoleService interface {
	ListRole(ctx context.Context, params model.ListRoleParams, filterParams types.FilterRoleParams) ([]model.Role, int64, error)
	GetRole(ctx context.Context, roleId int64) (model.Role, error)
	CreateRole(ctx context.Context, params model.CreateRoleParams) error
	UpdateRole(ctx context.Context, params model.UpdateRoleParams) error
	DeleteRole(ctx context.Context, roleId int64) error
}

// INoticeService 會員通知訊息
type INoticeService interface {
	ListNotice(ctx context.Context, params model.ListNoticeParams, filterParams types.FilterNoticeParams) ([]model.Notice, int64, error)
	GetNotice(ctx context.Context, noticeId int64) (model.Notice, error)
	CreateNotice(ctx context.Context, params model.CreateNoticeParams) error
	UpdateNotice(ctx context.Context, params model.UpdateNoticeParams) error
	DeleteNotice(ctx context.Context, noticeId int64) error
	GetLatestNotice(ctx context.Context) (model.Notice, error)
}

type IRemindService interface {
	ListRemind(ctx context.Context, params model.ListRemindParams, filterParams types.FilterRemindParams) ([]model.Remind, int64, error)
	GetRemind(ctx context.Context, remindId int64) (model.Remind, error)
	CreateRemind(ctx context.Context, params model.CreateRemindParams) error
	UpdateRemind(ctx context.Context, params model.UpdateRemindParams) error
	DeleteRemind(ctx context.Context, remindId int64) error
}
