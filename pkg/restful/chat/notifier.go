package chat

import (
	"context"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/model"
	"cs-api/pkg/types"
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
		OpType: types.OpTypeMemberJoin,
		Payload: map[string]interface{}{
			"room_id":     member.RoomID,
			"member_name": member.Name,
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, staff)
}

// StaffTyping 通知用戶客服正在輸入
func (n *Notifier) StaffTyping(staffName string, member *MemberClient) {
	if member == nil {
		return
	}

	message := model.Message{
		OpType: types.OpTypeStaffTyping,
		Payload: map[string]interface{}{
			"room_id":    member.RoomID,
			"staff_name": staffName,
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, member)
}

// Broadcast 廣播消息
func (n *Notifier) Broadcast(clientMessage ClientMessage, clients ...Client) {
	message := model.Message{
		OpType: types.OpTypeMessageReceived,
		Payload: map[string]interface{}{
			"room_id":      clientMessage.RoomID,
			"sender_type":  clients[0].GetSenderType(),
			"sender_id":    clients[0].GetID(),
			"sender_name":  clients[0].GetName(),
			"content_type": clientMessage.ContentType,
			"content":      clientMessage.Content,
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, true, clients...)
}

// NoStaff 通知客戶當前沒有客服可以服務
func (n *Notifier) NoStaff(member *MemberClient) {
	message := model.Message{
		OpType: types.OpTypeNoStaff,
		Payload: map[string]interface{}{
			"room_id": member.RoomID,
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, member)
}

// RoomClosed 通知客戶房間已關閉
func (n *Notifier) RoomClosed(member *MemberClient) {
	message := model.Message{
		OpType: types.OpTypeRoomClosed,
		Payload: map[string]interface{}{
			"room_id": member.RoomID,
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, member)
}

// SendScore 通知用戶評分
func (n *Notifier) SendScore(roomId int64, clients ...Client) {
	message := model.Message{
		OpType: types.OpTypeSendScore,
		Payload: map[string]interface{}{
			"room_id":     roomId,
			"sender_type": types.SenderTypeSystem,
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, true, clients...)
}

// MemberScored 通知客服用戶已完成評分
func (n *Notifier) MemberScored(member *MemberClient, staff *StaffClient) {
	message := model.Message{
		OpType: types.OpTypeSendScore,
		Payload: map[string]interface{}{
			"room_id":      member.RoomID,
			"sender_type":  types.SenderTypeSystem,
			"content_type": types.ContentTypeText,
			"content":      member.Name + " 已完成評分",
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, true, staff)
}

// RoomAccepted 通知客戶客服已接受並進入房間
func (n *Notifier) RoomAccepted(member *MemberClient) {
	message := model.Message{
		OpType: types.OpTypeRoomAccepted,
		Payload: map[string]interface{}{
			"room_id": member.RoomID,
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, member)
}

// Greeting 通知客戶問候語
func (n *Notifier) Greeting(content string, member *MemberClient, staff *StaffClient) {
	if content == "" {
		return
	}

	message := model.Message{
		OpType: types.OpTypeMessageReceived,
		Payload: map[string]interface{}{
			"room_id":      member.RoomID,
			"sender_type":  types.SenderTypeStaff,
			"sender_id":    staff.ID,
			"sender_name":  staff.Name,
			"content_type": types.ContentTypeText,
			"content":      content,
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, member)
}

// RoomTransferred 通知用戶已被轉接
func (n *Notifier) RoomTransferred(member *MemberClient) {
	message := model.Message{
		OpType: types.OpTypeRoomTransferred,
		Payload: map[string]interface{}{
			"room_id": member.RoomID,
		},
		Timestamp: time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, member)
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