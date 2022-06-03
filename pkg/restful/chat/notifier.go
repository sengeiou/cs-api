package chat

import (
	"context"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/model"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"reflect"
	"time"
)

type Notifier struct {
	msgSvc iface.IMessageService
}

func NewNotifier(msgSvc iface.IMessageService) *Notifier {
	return &Notifier{
		msgSvc: msgSvc,
	}
}

func (n *Notifier) MemberJoin(member *MemberClient, staff *StaffClient) {
	message := model.Message{
		Type:        model.MessageTypeSystem,
		RoomID:      member.RoomID,
		SenderName:  "系統",
		ContentType: model.MessageContentTypeJoin,
		Content:     member.Name + " 已進入房間",
		ExtraInfo: &model.ExtraInfo{
			ClientName: &member.Name,
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, true, staff)
}

func (n *Notifier) MemberLeave(member *MemberClient, staff *StaffClient) {
	message := model.Message{
		Type:        model.MessageTypeSystem,
		RoomID:      member.RoomID,
		SenderName:  "系統",
		ContentType: model.MessageContentTypeLeave,
		Content:     member.Name + " 已離開房間",
		Timestamp:   time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, true, staff)
}

// Typing 通知用戶客服正在輸入
func (n *Notifier) Typing(staffName string, member *MemberClient) {
	if member == nil {
		return
	}

	message := model.Message{
		Type:        model.MessageTypeSystem,
		RoomID:      member.RoomID,
		SenderName:  "系統",
		ContentType: model.MessageContentTypeTyping,
		Content:     staffName + " 正在輸入",
		Timestamp:   time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, false, member)
}

// Broadcast 廣播消息
func (n *Notifier) Broadcast(clientMessage ClientMessage, clients ...Client) {
	message := model.Message{
		Type:        clients[0].GetMessageType(),
		RoomID:      clientMessage.RoomID,
		SenderID:    clients[0].GetID(),
		SenderName:  clients[0].GetName(),
		ContentType: clientMessage.ContentType,
		Content:     clientMessage.Content,
		Timestamp:   time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, true, clients...)
}

// NoStaff 通知客戶當前沒有客服可以服務
func (n *Notifier) NoStaff(member *MemberClient) {
	message := model.Message{
		Type:        model.MessageTypeSystem,
		RoomID:      member.RoomID,
		SenderName:  "系統",
		ContentType: model.MessageContentTypeNoStaff,
		Content:     "客服人員忙線中，請稍候在試",
		Timestamp:   time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, false, member)
}

// RoomClosed 通知客戶房間已關閉
func (n *Notifier) RoomClosed(member *MemberClient) {
	message := model.Message{
		Type:        model.MessageTypeSystem,
		RoomID:      member.RoomID,
		SenderName:  "系統",
		ContentType: model.MessageContentTypeRoomClosed,
		Content:     "諮詢已結束",
		Timestamp:   time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, true, member)
}

// SendScore 通知用戶評分
func (n *Notifier) SendScore(roomId int64, clients ...Client) {
	message := model.Message{
		Type:        model.MessageTypeSystem,
		RoomID:      roomId,
		SenderName:  "系統",
		ContentType: model.MessageContentTypeScore,
		Content:     "已發送評分請求",
		Timestamp:   time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, true, clients...)
}

// MemberScored 通知客服用戶已完成評分
func (n *Notifier) MemberScored(member *MemberClient, staff *StaffClient) {
	message := model.Message{
		Type:        model.MessageTypeSystem,
		RoomID:      member.RoomID,
		SenderName:  "系統",
		ContentType: model.MessageContentTypeScore,
		Content:     member.Name + " 已完成評分",
		Timestamp:   time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, true, staff)
}

// RoomAccepted 通知客戶客服已接受並進入房間
func (n *Notifier) RoomAccepted(member *MemberClient) {
	message := model.Message{
		Type:        model.MessageTypeSystem,
		RoomID:      member.RoomID,
		SenderName:  "系統",
		ContentType: model.MessageContentTypeRoomAccepted,
		Content:     "",
		Timestamp:   time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, false, member)
}

// Greeting 通知客戶問候語
func (n *Notifier) Greeting(content string, member *MemberClient, staff *StaffClient) {
	if content == "" {
		return
	}

	message := model.Message{
		Type:        model.MessageTypeStaff,
		RoomID:      member.RoomID,
		SenderID:    staff.ID,
		SenderName:  staff.Name,
		ContentType: model.MessageContentTypeText,
		Content:     content,
		Timestamp:   time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, false, member)
}

// RoomTransferred 通知用戶已被轉接
func (n *Notifier) RoomTransferred(member *MemberClient) {
	message := model.Message{
		Type:        model.MessageTypeSystem,
		RoomID:      member.RoomID,
		SenderName:  "系統",
		ContentType: model.MessageContentTypeText,
		Content:     "您已被轉接",
		Timestamp:   time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, true, member)
}

func (n *Notifier) send(message model.Message, needStore bool, clients ...Client) {
	if needStore {
		if err := n.msgSvc.CreateMessage(context.Background(), message); err != nil {
			log.Error().Msgf("create message error: %s\n", err)
		}
	}

	byteMessage, _ := json.Marshal(message)
	for _, client := range clients {
		if reflect.ValueOf(client).IsNil() {
			continue
		}
		ch := client.GetSendChan()
		ch <- byteMessage
	}
}
