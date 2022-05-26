// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package model

import (
	"context"
	"database/sql"
	"time"

	"cs-api/pkg/types"
)

type Querier interface {
	AcceptRoom(ctx context.Context, arg AcceptRoomParams) error
	CloseRoom(ctx context.Context, arg CloseRoomParams) error
	ConstantSeeder(ctx context.Context) error
	// 計算每日分類諮詢數
	CountClosedRoomByTag(ctx context.Context, arg CountClosedRoomByTagParams) ([]CountClosedRoomByTagRow, error)
	// 計算每日訪客數
	CountDailyRoomByMember(ctx context.Context, arg CountDailyRoomByMemberParams) (int64, error)
	CountListFastMessage(ctx context.Context) (int64, error)
	CountListNotice(ctx context.Context) (int64, error)
	CountListRemind(ctx context.Context) (int64, error)
	CountListRole(ctx context.Context) (int64, error)
	CountListRoom(ctx context.Context) (int64, error)
	CountListStaff(ctx context.Context) (int64, error)
	CountListStaffRoom(ctx context.Context, status types.RoomStatus) (int64, error)
	CountListTag(ctx context.Context) (int64, error)
	CreateFastMessage(ctx context.Context, arg CreateFastMessageParams) error
	CreateFastMessageCategory(ctx context.Context, arg CreateFastMessageCategoryParams) error
	CreateMember(ctx context.Context, arg CreateMemberParams) (sql.Result, error)
	CreateNotice(ctx context.Context, arg CreateNoticeParams) error
	CreateRemind(ctx context.Context, arg CreateRemindParams) error
	CreateReportDailyGuest(ctx context.Context, arg CreateReportDailyGuestParams) error
	CreateReportDailyTag(ctx context.Context, arg CreateReportDailyTagParams) error
	CreateRole(ctx context.Context, arg CreateRoleParams) error
	CreateRoom(ctx context.Context, arg CreateRoomParams) (sql.Result, error)
	CreateStaff(ctx context.Context, arg CreateStaffParams) error
	CreateTag(ctx context.Context, arg CreateTagParams) error
	DeleteFastMessage(ctx context.Context, id int64) error
	DeleteNotice(ctx context.Context, id int64) error
	DeleteRemind(ctx context.Context, id int64) error
	DeleteReportDailyGuest(ctx context.Context, date time.Time) error
	DeleteReportDailyTag(ctx context.Context, date time.Time) error
	DeleteRole(ctx context.Context, id int64) error
	DeleteStaff(ctx context.Context, id int64) error
	DeleteTag(ctx context.Context, id int64) error
	GetAllAvailableFastMessage(ctx context.Context) ([]GetAllAvailableFastMessageRow, error)
	GetAllTag(ctx context.Context) ([]Tag, error)
	GetAvailableNotice(ctx context.Context) (Notice, error)
	GetCsConfig(ctx context.Context) (Constant, error)
	GetFastMessage(ctx context.Context, id int64) (FastMessage, error)
	GetGuestMember(ctx context.Context, deviceID string) (Member, error)
	// 獲取會員並未關閉的房間
	GetMemberAvailableRoom(ctx context.Context, memberID int64) (Room, error)
	GetNormalMember(ctx context.Context, name string) (Member, error)
	GetNotice(ctx context.Context, id int64) (Notice, error)
	GetRemind(ctx context.Context, id int64) (Remind, error)
	GetRole(ctx context.Context, id int64) (Role, error)
	GetRoom(ctx context.Context, id int64) (GetRoomRow, error)
	GetStaff(ctx context.Context, id int64) (GetStaffRow, error)
	GetStaffCountByRoleId(ctx context.Context, roleID int64) (int64, error)
	GetStaffRoom(ctx context.Context, staffID int64) ([]int64, error)
	GetTag(ctx context.Context, id int64) (GetTagRow, error)
	ListAvailableStaff(ctx context.Context, id int64) ([]Staff, error)
	ListFastMessage(ctx context.Context, arg ListFastMessageParams) ([]ListFastMessageRow, error)
	ListFastMessageCategory(ctx context.Context) ([]Constant, error)
	ListNotice(ctx context.Context, arg ListNoticeParams) ([]Notice, error)
	ListRemind(ctx context.Context, arg ListRemindParams) ([]Remind, error)
	ListReportDailyGuest(ctx context.Context, arg ListReportDailyGuestParams) ([]ReportDailyGuest, error)
	ListReportDailyTag(ctx context.Context, arg ListReportDailyTagParams) ([]ReportDailyTag, error)
	ListRole(ctx context.Context, arg ListRoleParams) ([]Role, error)
	ListRoom(ctx context.Context, arg ListRoomParams) ([]ListRoomRow, error)
	ListStaff(ctx context.Context, arg ListStaffParams) ([]ListStaffRow, error)
	ListStaffRoom(ctx context.Context, arg ListStaffRoomParams) ([]ListStaffRoomRow, error)
	ListTag(ctx context.Context, arg ListTagParams) ([]ListTagRow, error)
	RoleSeeder(ctx context.Context) error
	StaffLogin(ctx context.Context, arg StaffLoginParams) (Staff, error)
	StaffSeeder(ctx context.Context, password string) error
	TagSeeder(ctx context.Context) error
	UpdateCsConfig(ctx context.Context, arg UpdateCsConfigParams) error
	UpdateFastMessage(ctx context.Context, arg UpdateFastMessageParams) error
	UpdateNotice(ctx context.Context, arg UpdateNoticeParams) error
	UpdateRemind(ctx context.Context, arg UpdateRemindParams) error
	UpdateRole(ctx context.Context, arg UpdateRoleParams) error
	UpdateRoomScore(ctx context.Context, arg UpdateRoomScoreParams) error
	UpdateRoomStaff(ctx context.Context, arg UpdateRoomStaffParams) error
	UpdateStaff(ctx context.Context, arg UpdateStaffParams) error
	UpdateStaffAvatar(ctx context.Context, arg UpdateStaffAvatarParams) error
	UpdateStaffLogin(ctx context.Context, arg UpdateStaffLoginParams) error
	UpdateStaffServingStatus(ctx context.Context, arg UpdateStaffServingStatusParams) error
	UpdateStaffWithPassword(ctx context.Context, arg UpdateStaffWithPasswordParams) error
	UpdateTag(ctx context.Context, arg UpdateTagParams) error
}

var _ Querier = (*Queries)(nil)
