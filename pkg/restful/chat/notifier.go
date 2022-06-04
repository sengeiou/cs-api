package chat

import (
	"context"
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
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
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeMemberJoin,
		Content:   member.Name,
		Ts:        time.Now().Unix(),
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
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeStaffTyping,
		Content:   staffName,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, member)
}

// Broadcast 廣播消息
func (n *Notifier) Broadcast(clientMessage ClientMessage, clients ...Client) {
	message := model.Message{
		RoomID:      clientMessage.RoomID,
		OpType:      types.MessageOpTypeMessageReceived,
		SenderType:  clients[0].GetSenderType(),
		SenderID:    clients[0].GetID(),
		SenderName:  clients[0].GetName(),
		ContentType: clientMessage.ContentType,
		Content:     clientMessage.Content,
		Extra:       json.RawMessage(`{"a":1}`),
		Ts:          time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, true, clients...)
}

// NoStaff 通知客戶當前沒有客服可以服務
func (n *Notifier) NoStaff(member *MemberClient) {
	message := model.Message{
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeNoStaff,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, member)
}

// RoomClosed 通知客戶房間已關閉
func (n *Notifier) RoomClosed(member *MemberClient) {
	message := model.Message{
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeRoomClosed,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, member)
}

// SendScore 通知用戶評分
func (n *Notifier) SendScore(roomId int64, clients ...Client) {
	message := model.Message{
		RoomID:    roomId,
		OpType:    types.MessageOpTypeSendScore,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, true, clients...)
}

// MemberScored 通知客服用戶已完成評分
func (n *Notifier) MemberScored(member *MemberClient, staff *StaffClient) {
	message := model.Message{
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeCompleteScore,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, true, staff)
}

// RoomAccepted 通知客戶客服已接受並進入房間
func (n *Notifier) RoomAccepted(member *MemberClient) {
	message := model.Message{
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeRoomAccepted,
		Ts:        time.Now().Unix(),
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
		RoomID:      member.RoomID,
		OpType:      types.MessageOpTypeMessageReceived,
		SenderType:  types.MessageSenderTypeStaff,
		SenderID:    staff.ID,
		SenderName:  staff.Name,
		ContentType: types.MessageContentTypeText,
		Content:     content,
		Ts:          time.Now().Unix(),
		CreatedAt:   time.Now().UTC(),
	}

	n.send(message, false, member)
}

// RoomTransferred 通知用戶已被轉接
func (n *Notifier) RoomTransferred(member *MemberClient) {
	message := model.Message{
		RoomID:    member.RoomID,
		OpType:    types.MessageOpTypeRoomTransferred,
		Ts:        time.Now().Unix(),
		CreatedAt: time.Now().UTC(),
	}

	n.send(message, false, member)
}

func (n *Notifier) send(message model.Message, needStore bool, clients ...Client) {
	if needStore {
		params := model.CreateMessageParams{
			RoomID:      message.RoomID,
			OpType:      message.OpType,
			SenderType:  message.SenderType,
			SenderID:    message.SenderID,
			SenderName:  message.SenderName,
			ContentType: message.ContentType,
			Content:     message.Content,
			Extra:       message.Extra,
			Ts:          message.Ts,
			CreatedAt:   message.CreatedAt,
		}
		if params.Extra == nil {
			params.Extra = json.RawMessage(`{}`)
		}
		if err := n.msgSvc.CreateMessage(context.Background(), params); err != nil {
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
