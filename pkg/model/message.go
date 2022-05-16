package model

import (
	"time"
)

type MessageType int8

const (
	MessageTypeSystem MessageType = iota + 1
	MessageTypeMember
	MessageTypeStaff
)

type MessageContentType int8

const (
	// MessageContentTypeTyping 輸入中
	MessageContentTypeTyping MessageContentType = iota + 1
	// MessageContentTypeText 輸入文字
	MessageContentTypeText
	// MessageContentTypeImage 上傳圖片
	MessageContentTypeImage
	// MessageContentTypeScore 發送評分請求
	MessageContentTypeScore
	// MessageContentTypeJoin 加入房間
	MessageContentTypeJoin
	// MessageContentTypeLeave 離開房間
	MessageContentTypeLeave
	// MessageContentTypeNoStaff 無客服可以服務客戶
	MessageContentTypeNoStaff
	// MessageContentTypeRoomClosed 關閉諮詢房
	MessageContentTypeRoomClosed
	// MessageContentTypeRoomAccepted 客服開始諮詢等待中諮詢房
	MessageContentTypeRoomAccepted
)

type ExtraInfo struct {
	ClientName *string `bson:"client_name,omitempty" json:"client_name"`
}

type Message struct {
	ID          string             `bson:"_id,omitempty" json:"id,omitempty"`
	Type        MessageType        `bson:"type" json:"type,omitempty"`
	RoomID      int64              `bson:"room_id" json:"room_id,omitempty"`
	SenderID    int64              `bson:"sender_id" json:"sender_id,omitempty"`
	SenderName  string             `bson:"sender_name" json:"sender_name,omitempty"`
	ContentType MessageContentType `bson:"content_type" json:"content_type,omitempty"`
	Content     string             `bson:"content" json:"content,omitempty"`
	ExtraInfo   *ExtraInfo         `bson:"extra_info,omitempty" json:"extra_info,omitempty"`
	Timestamp   int64              `bson:"timestamp" json:"timestamp,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at,omitempty"`
}

func (Message) GetCollection() string {
	return "messages"
}
