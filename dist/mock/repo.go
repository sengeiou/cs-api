package mock

import (
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func NewRepository(t *testing.T) iface.IRepository {
	m := gomock.NewController(t)
	mock := NewMockIRepository(m)
	mock.EXPECT().WithTx(gomock.Any()).AnyTimes().Return((*model.Queries)(nil))
	mock.EXPECT().Transaction(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	// Tag
	mock.EXPECT().GetTag(gomock.Any(), gomock.Any()).AnyTimes().Return(model.GetTagRow{}, nil)
	mock.EXPECT().CreateTag(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateTag(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().DeleteTag(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().ListAvailableTag(gomock.Any()).AnyTimes().Return(make([]model.ListAvailableTagRow, 0), nil)

	// Role
	mock.EXPECT().GetRole(gomock.Any(), gomock.Any()).AnyTimes().Return(model.Role{}, nil)
	mock.EXPECT().CreateRole(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateRole(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().DeleteRole(gomock.Any(), int64(1)).AnyTimes().Return(errors.New("unable to delete"))
	mock.EXPECT().DeleteRole(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	// Staff
	mock.EXPECT().GetStaff(gomock.Any(), int64(1)).AnyTimes().Return(model.GetStaffRow{ServingStatus: types.StaffServingStatusClosed}, nil)
	mock.EXPECT().GetStaff(gomock.Any(), int64(2)).AnyTimes().Return(model.GetStaffRow{ServingStatus: types.StaffServingStatusServing}, nil)
	mock.EXPECT().GetStaff(gomock.Any(), gomock.Any()).AnyTimes().Return(model.GetStaffRow{}, nil)
	mock.EXPECT().CreateStaff(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateStaff(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateStaffWithPassword(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateStaffAvatar(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().DeleteStaff(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().GetStaffCountByRoleId(gomock.Any(), int64(1)).AnyTimes().Return(int64(1), nil)
	mock.EXPECT().GetStaffCountByRoleId(gomock.Any(), gomock.Any()).AnyTimes().Return(int64(0), nil)
	mock.EXPECT().UpdateStaffServingStatus(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().ListAvailableStaff(gomock.Any(), gomock.Any()).AnyTimes().Return(make([]model.ListAvailableStaffRow, 0), nil)
	mock.EXPECT().StaffLogin(gomock.Any(), gomock.Any()).AnyTimes().Return(model.StaffLoginRow{ID: 1, Name: "user", Username: "user", Permissions: json.RawMessage(`["*"]`)}, nil)
	mock.EXPECT().UpdateStaffLogin(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	// Member
	mock.EXPECT().GetGuestMember(gomock.Any(), "deviceId").AnyTimes().Return(model.Member{ID: 1}, nil)
	mock.EXPECT().GetGuestMember(gomock.Any(), gomock.Any()).AnyTimes().Return(model.Member{}, sql.ErrNoRows)
	mock.EXPECT().GetNormalMember(gomock.Any(), "name").AnyTimes().Return(model.Member{ID: 1}, nil)
	mock.EXPECT().GetNormalMember(gomock.Any(), gomock.Any()).AnyTimes().Return(model.Member{}, sql.ErrNoRows)
	mock.EXPECT().CreateMember(gomock.Any(), gomock.Any()).AnyTimes().Return(MockSqlResult{}, nil)
	mock.EXPECT().GetOnlineStatus(gomock.Any(), gomock.Any()).AnyTimes().Return(types.MemberOnlineStatus(1), nil)
	mock.EXPECT().UpdateOnlineStatus(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	// Room
	mock.EXPECT().GetMemberAvailableRoom(gomock.Any(), int64(1)).AnyTimes().Return(model.Room{ID: 1, MemberID: 1}, nil)
	mock.EXPECT().GetMemberAvailableRoom(gomock.Any(), gomock.Any()).AnyTimes().Return(model.Room{}, sql.ErrNoRows)
	mock.EXPECT().CreateRoom(gomock.Any(), gomock.Any()).AnyTimes().Return(MockSqlResult{}, nil)
	mock.EXPECT().GetRoom(gomock.Any(), int64(1)).AnyTimes().Return(model.GetRoomRow{ID: 1, Status: types.RoomStatusPending}, nil)
	mock.EXPECT().GetRoom(gomock.Any(), int64(2)).AnyTimes().Return(model.GetRoomRow{ID: 2, StaffID: 1, Status: types.RoomStatusServing}, nil)
	mock.EXPECT().GetRoom(gomock.Any(), int64(3)).AnyTimes().Return(model.GetRoomRow{ID: 3, StaffID: 2, Status: types.RoomStatusServing}, nil)
	mock.EXPECT().GetRoom(gomock.Any(), gomock.Any()).AnyTimes().Return(model.GetRoomRow{ID: 4, Status: types.RoomStatusClosed}, nil)
	mock.EXPECT().AcceptRoom(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().CloseRoom(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateRoomScore(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().GetStaffRoom(gomock.Any(), gomock.Any()).AnyTimes().Return(make([]int64, 0), nil)
	mock.EXPECT().UpdateRoomStaff(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	// Constant
	mock.EXPECT().GetCsConfig(gomock.Any()).AnyTimes().Return(model.Constant{Value: `{"max_member":5}`}, nil)
	mock.EXPECT().UpdateCsConfig(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().ListFastReplyCategory(gomock.Any()).AnyTimes().Return(make([]model.Constant, 0), nil)
	mock.EXPECT().CreateFastReplyCategory(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().CheckFastReplyCategory(gomock.Any(), gomock.Any()).AnyTimes().Return(1, nil)

	// FastReply
	mock.EXPECT().GetFastReply(gomock.Any(), gomock.Any()).AnyTimes().Return(model.FastReply{}, nil)
	mock.EXPECT().CreateFastReply(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateFastReply(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().DeleteFastReply(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().GetAllAvailableFastReply(gomock.Any()).AnyTimes().Return([]model.GetAllAvailableFastReplyRow{
		{
			CategoryID: 1,
			Title:      "測試1",
			Content:    "測試內容1",
			Category:   "分類1",
		},
	}, nil)

	// Notice
	mock.EXPECT().GetNotice(gomock.Any(), gomock.Any()).AnyTimes().Return(model.Notice{}, nil)
	mock.EXPECT().CreateNotice(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateNotice(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().DeleteNotice(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().GetLatestNotice(gomock.Any()).AnyTimes().Return(model.GetLatestNoticeRow{}, nil)

	// Remind
	mock.EXPECT().GetRemind(gomock.Any(), gomock.Any()).AnyTimes().Return(model.Remind{}, nil)
	mock.EXPECT().CreateRemind(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateRemind(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().DeleteRemind(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().ListActiveRemind(gomock.Any()).AnyTimes().Return(make([]model.ListActiveRemindRow, 0), nil)

	// FAQ
	mock.EXPECT().GetFAQ(gomock.Any(), gomock.Any()).AnyTimes().Return(model.GetFAQRow{}, nil)
	mock.EXPECT().CreateFAQ(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateFAQ(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().DeleteFAQ(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().ListAvailableFAQ(gomock.Any()).AnyTimes().Return(make([]model.ListAvailableFAQRow, 0), nil)

	// Message
	mock.EXPECT().CreateMessage(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().ListMemberRoomMessage(gomock.Any(), gomock.Any()).AnyTimes().Return(make([]model.Message, 0), nil)
	mock.EXPECT().ListStaffRoomMessage(gomock.Any(), gomock.Any()).AnyTimes().Return(make([]model.Message, 0), nil)

	// Merchant
	mock.EXPECT().GetMerchant(gomock.Any(), gomock.Any()).AnyTimes().Return(model.GetMerchantRow{}, nil)
	mock.EXPECT().CreateMerchant(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UpdateMerchant(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().DeleteMerchant(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().ListAvailableMerchant(gomock.Any()).AnyTimes().Return(make([]model.ListAvailableMerchantRow, 0), nil)
	mock.EXPECT().CheckMerchantKey(gomock.Any(), gomock.Any()).AnyTimes().Return(int64(0), nil)

	return mock
}
