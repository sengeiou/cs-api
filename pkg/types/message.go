package types

type OpType int8

const (
	// OpTypeStaffTyping 客服輸入中
	OpTypeStaffTyping OpType = iota + 1
	// OpTypeMessageReceived 收到聊天訊息
	OpTypeMessageReceived
	// OpTypeSendScore 發送評分請求
	OpTypeSendScore
	// OpTypeMemberJoin 會員加入房間
	OpTypeMemberJoin
	// OpTypeNoStaff 無客服可以服務客戶
	OpTypeNoStaff
	// OpTypeRoomClosed 關閉諮詢房
	OpTypeRoomClosed
	// OpTypeRoomAccepted 客服開始諮詢等待中諮詢房
	OpTypeRoomAccepted
	// OpTypeRoomTransferred 轉接諮詢房
	OpTypeRoomTransferred
)

type ContentType int8

const (
	ContentTypeText ContentType = iota + 1
	ContentTypeImage
)

type SenderType int8

const (
	SenderTypeSystem SenderType = iota + 1
	SenderTypeMember
	SenderTypeStaff
)

type ListMessageParams struct {
	RoomID   int64
	StaffID  int64
	Content  string
	Page     int64
	PageSize int64
}
