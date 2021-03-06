package iface

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	"cs-api/pkg/types"
	"github.com/gin-gonic/gin"
	"time"
)

type IAuthService interface {
	Login(ctx context.Context, username, password string) (pkg.ClientInfo, error)
	GetStaffInfo(ctx context.Context, staffId int64) (model.GetStaffRow, error)
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
	ListAvailableStaff(ctx context.Context, staffId int64) ([]model.ListAvailableStaffRow, error)
	GetAllStaffs(ctx context.Context) ([]model.GetAllStaffsRow, error)
}

type IRoomService interface {
	CreateRoom(ctx context.Context, deviceId string, name string) (model.Room, model.Member, error)
	AcceptRoom(ctx context.Context, staffId int64, roomId int64) error
	CloseRoom(ctx context.Context, staffId, roomId, tagId int64) error
	TransferRoom(ctx context.Context, staffId, roomId, toStaffId int64) error
	UpdateRoomScore(ctx context.Context, roomId int64, score int32) error
	ListStaffRoom(ctx context.Context, params model.ListStaffRoomParams, filterParams types.FilterStaffRoomParams) ([]model.ListStaffRoomRow, int64, error)
	ListRoom(ctx context.Context, params model.ListRoomParams, filterParams types.FilterRoomParams) ([]types.Room, int64, error)
	GetStaffRooms(ctx context.Context, staffId int64) ([]int64, error)
}

type IMessageService interface {
	CreateMessage(ctx context.Context, params model.CreateMessageParams) error
	ListRoomMessage(ctx context.Context, params interface{}) ([]model.Message, error)
	ListMessage(ctx context.Context, params model.ListMessageParams, filterParams types.FilterMessageParams) ([]model.Message, int64, error)
}

type ITagService interface {
	ListTag(ctx context.Context, params model.ListTagParams, filterParams types.FilterTagParams) ([]model.ListTagRow, int64, error)
	GetTag(ctx context.Context, tagId int64) (model.GetTagRow, error)
	CreateTag(ctx context.Context, params model.CreateTagParams) error
	UpdateTag(ctx context.Context, params model.UpdateTagParams) error
	DeleteTag(ctx context.Context, tagId int64) error
	ListAvailableTag(ctx context.Context) ([]model.ListAvailableTagRow, error)
}

type IFastReplyService interface {
	ListFastReply(ctx context.Context, params model.ListFastReplyParams, filterParams types.FilterFastReplyParams) ([]model.ListFastReplyRow, int64, error)
	GetFastReply(ctx context.Context, id int64) (model.FastReply, error)
	CreateFastReply(ctx context.Context, params model.CreateFastReplyParams) error
	UpdateFastReply(ctx context.Context, params model.UpdateFastReplyParams) error
	DeleteFastReply(ctx context.Context, id int64) error
	ListCategory(ctx context.Context) ([]model.Constant, error)
	CreateCategory(ctx context.Context, params model.CreateFastReplyCategoryParams) error
	CheckCategory(ctx context.Context, id int64) (interface{}, error)
	ListFastReplyGroup(ctx context.Context) ([]pkg.FastReplyGroupItem, error)
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
	ListMember(ctx context.Context, params model.ListMemberParams, filterParams types.FilterMemberParams) ([]model.Member, int64, error)
	GetOrCreateMember(ctx context.Context, name string, deviceId string) (model.Member, error)
	GetOnlineStatus(ctx context.Context, memberId int64) (types.MemberOnlineStatus, error)
	UpdateOnlineStatus(ctx context.Context, params model.UpdateOnlineStatusParams) error
}

type IRoleService interface {
	ListRole(ctx context.Context, params model.ListRoleParams, filterParams types.FilterRoleParams) ([]model.Role, int64, error)
	GetAllRoles(ctx context.Context) ([]model.GetAllRolesRow, error)
	GetRole(ctx context.Context, roleId int64) (model.Role, error)
	CreateRole(ctx context.Context, params model.CreateRoleParams) error
	UpdateRole(ctx context.Context, params model.UpdateRoleParams) error
	DeleteRole(ctx context.Context, roleId int64) error
}

// INoticeService ??????????????????
type INoticeService interface {
	ListNotice(ctx context.Context, params model.ListNoticeParams, filterParams types.FilterNoticeParams) ([]types.Notice, int64, error)
	GetNotice(ctx context.Context, noticeId int64) (model.Notice, error)
	CreateNotice(ctx context.Context, params model.CreateNoticeParams) error
	UpdateNotice(ctx context.Context, params model.UpdateNoticeParams) error
	DeleteNotice(ctx context.Context, noticeId int64) error
	GetLatestNotice(ctx context.Context) (model.GetLatestNoticeRow, error)
}

type IRemindService interface {
	ListRemind(ctx context.Context, params model.ListRemindParams, filterParams types.FilterRemindParams) ([]model.Remind, int64, error)
	GetRemind(ctx context.Context, remindId int64) (model.Remind, error)
	CreateRemind(ctx context.Context, params model.CreateRemindParams) error
	UpdateRemind(ctx context.Context, params model.UpdateRemindParams) error
	DeleteRemind(ctx context.Context, remindId int64) error
	ListActiveRemind(ctx context.Context) ([]model.ListActiveRemindRow, error)
}

type IFAQService interface {
	ListFAQ(ctx context.Context, params model.ListFAQParams, filterParams types.FilterFAQParams) ([]model.ListFAQRow, int64, error)
	GetFAQ(ctx context.Context, faqId int64) (model.GetFAQRow, error)
	CreateFAQ(ctx context.Context, params model.CreateFAQParams) error
	UpdateFAQ(ctx context.Context, params model.UpdateFAQParams) error
	DeleteFAQ(ctx context.Context, faqId int64) error
	ListAvailableFAQ(ctx context.Context) ([]model.ListAvailableFAQRow, error)
}

type IMerchantService interface {
	ListMerchant(ctx context.Context, params model.ListMerchantParams, filterParams types.FilterMerchantParams) ([]model.ListMerchantRow, int64, error)
	GetMerchant(ctx context.Context, merchantId int64) (model.GetMerchantRow, error)
	CreateMerchant(ctx context.Context, params model.CreateMerchantParams) error
	UpdateMerchant(ctx context.Context, params model.UpdateMerchantParams) error
	DeleteMerchant(ctx context.Context, merchantId int64) error
	ListAvailableMerchant(ctx context.Context) ([]model.ListAvailableMerchantRow, error)
}
