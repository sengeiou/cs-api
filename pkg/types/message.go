package types

type MessageOpType int8

const (
	// MessageOpTypeStaffTyping 客服輸入中
	MessageOpTypeStaffTyping MessageOpType = iota + 1
	// MessageOpTypeMessageReceived 收到聊天訊息
	MessageOpTypeMessageReceived
	// MessageOpTypeSendScore 發送評分請求
	MessageOpTypeSendScore
	// MessageOpTypeCompleteScore 用戶完成評分
	MessageOpTypeCompleteScore
	// MessageOpTypeMemberJoin 會員加入房間
	MessageOpTypeMemberJoin
	// MessageOpTypeNoStaff 無客服可以服務客戶
	MessageOpTypeNoStaff
	// MessageOpTypeRoomClosed 關閉諮詢房
	MessageOpTypeRoomClosed
	// MessageOpTypeRoomAccepted 客服開始諮詢等待中諮詢房
	MessageOpTypeRoomAccepted
	// MessageOpTypeRoomTransferred 轉接諮詢房
	MessageOpTypeRoomTransferred
)

type MessageSenderType int8

const (
	MessageSenderTypeSystem MessageSenderType = iota
	MessageSenderTypeMember
	MessageSenderTypeStaff
)

type MessageContentType int8

const (
	MessageContentTypeElse MessageContentType = iota
	MessageContentTypeText
	MessageContentTypeImage
)

type FilterMessageParams struct {
	RoomID  *int64
	StaffID *int64
	Content *string
}

type ListMessageParams struct {
	RoomID   int64
	StaffID  int64
	Content  string
	Page     int64
	PageSize int64
}
