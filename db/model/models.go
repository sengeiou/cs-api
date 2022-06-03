// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package model

import (
	"database/sql"
	"encoding/json"
	"time"

	"cs-api/pkg/types"
)

type Constant struct {
	ID int64 `db:"id" json:"id"`
	// 常數類型 1快捷回覆 2客服配置
	Type types.ConstantType `db:"type" json:"type"`
	// 鍵
	Key types.ConstantKey `db:"key" json:"key"`
	// 值
	Value string `db:"value" json:"value"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type FastReply struct {
	ID int64 `db:"id" json:"id"`
	// 分類ID(constantID)
	CategoryID int64 `db:"category_id" json:"category_id"`
	// 訊息標題
	Title string `db:"title" json:"title"`
	// 訊息內容
	Content string `db:"content" json:"content"`
	// 消息狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Member struct {
	ID int64 `db:"id" json:"id"`
	// 用戶類型 1一般用戶 2訪客
	Type types.MemberType `db:"type" json:"type"`
	// 會員名稱
	Name string `db:"name" json:"name"`
	// 設備號
	DeviceID string `db:"device_id" json:"device_id"`
	// 會員狀態 1在線 2離線
	Status types.MemberStatus `db:"status" json:"status"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 系統公告資料表
type Notice struct {
	ID int64 `db:"id" json:"id"`
	// 公告標題
	Title string `db:"title" json:"title"`
	// 公告內容
	Content string `db:"content" json:"content"`
	// 開始時間
	StartAt time.Time `db:"start_at" json:"start_at"`
	// 結束時間
	EndAt time.Time `db:"end_at" json:"end_at"`
	// 狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// 後台提醒事項資料表
type Remind struct {
	ID int64 `db:"id" json:"id"`
	// 標題
	Title string `db:"title" json:"title"`
	// 內容
	Content string `db:"content" json:"content"`
	// 狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type ReportDailyGuest struct {
	ID int64 `db:"id" json:"id"`
	// 報表日期
	Date time.Time `db:"date" json:"date"`
	// 訪客數
	GuestCount int32 `db:"guest_count" json:"guest_count"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type ReportDailyTag struct {
	ID int64 `db:"id" json:"id"`
	// 報表日期
	Date time.Time `db:"date" json:"date"`
	// 標籤ID
	TagID int64 `db:"tag_id" json:"tag_id"`
	// 人數
	Count int32 `db:"count" json:"count"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Role struct {
	ID int64 `db:"id" json:"id"`
	// 角色名稱
	Name string `db:"name" json:"name"`
	// 角色權限
	Permissions json.RawMessage `db:"permissions" json:"permissions"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Room struct {
	ID int64 `db:"id" json:"id"`
	// 職員ID
	StaffID int64 `db:"staff_id" json:"staff_id"`
	// 會員ID
	MemberID int64 `db:"member_id" json:"member_id"`
	// 標籤ID
	TagID int64 `db:"tag_id" json:"tag_id"`
	// 評分 1-5
	Score int32 `db:"score" json:"score"`
	// 客服房狀態 1等待中 2服務中 3已關閉
	Status types.RoomStatus `db:"status" json:"status"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	// 關閉時間
	ClosedAt sql.NullTime `db:"closed_at" json:"closed_at"`
}

type Staff struct {
	ID int64 `db:"id" json:"id"`
	// 角色ID
	RoleID int64 `db:"role_id" json:"role_id"`
	// 職員姓名
	Name string `db:"name" json:"name"`
	// 用戶名
	Username string `db:"username" json:"username"`
	// 密碼
	Password string `db:"password" json:"password"`
	// 大頭貼
	Avatar string `db:"avatar" json:"avatar"`
	// 職員狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 職員服務狀態 1關閉 2服務中 3閒置
	ServingStatus types.StaffServingStatus `db:"serving_status" json:"serving_status"`
	// 上次登入時間
	LastLoginTime sql.NullTime `db:"last_login_time" json:"last_login_time"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Tag struct {
	ID int64 `db:"id" json:"id"`
	// 標籤名稱
	Name string `db:"name" json:"name"`
	// 標籤狀態 1開啟 2關閉
	Status types.Status `db:"status" json:"status"`
	// 創建管理員
	CreatedBy int64 `db:"created_by" json:"created_by"`
	// 創建時間
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// 更新管理員
	UpdatedBy int64 `db:"updated_by" json:"updated_by"`
	// 更新時間
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
