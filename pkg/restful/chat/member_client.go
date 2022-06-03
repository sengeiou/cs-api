package chat

import (
	"cs-api/pkg"
	iface2 "cs-api/pkg/interface"
	"cs-api/pkg/model"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type MemberClient struct {
	ID         int64
	Type       pkg.ClientType
	Name       string
	RoomID     int64
	StaffID    int64 // 客服 ID
	Socket     *websocket.Conn
	SendChan   chan []byte
	Status     ClientStatus
	Manager    *ClientManager
	dispatcher *StaffDispatcher
	Notifier   *Notifier
	MsgSvc     iface2.IMessageService
	isSend     bool          // 是否已經發送過訊息
	Sending    chan struct{} // 為了提早取消檢查閒置時間的 goroutine
}

func (mc *MemberClient) GetID() int64 {
	return mc.ID
}

func (mc *MemberClient) GetName() string {
	return mc.Name
}

func (mc *MemberClient) GetType() pkg.ClientType {
	return mc.Type
}

func (mc *MemberClient) GetMessageType() model.MessageType {
	return model.MessageTypeMember
}

func (mc *MemberClient) GetStatus() ClientStatus {
	return mc.Status
}

func (mc *MemberClient) GetSendChan() chan []byte {
	return mc.SendChan
}

func (mc *MemberClient) SocketRead() {
	defer func() {
		mc.Manager.unregister <- mc
	}()

	for {
		_, message, err := mc.Socket.ReadMessage()
		if err != nil {
			log.Error().Msgf("member client ws error: %s\n", err)
			break
		} else {
			if !mc.isSend {
				mc.Sending <- struct{}{}
				mc.isSend = true
			}
		}

		var tmp ClientMessage
		if err = json.Unmarshal(message, &tmp); err != nil {
			log.Error().Msgf("error: %s\n", err)
			break
		}

		staff := mc.dispatcher.getStaff(mc.StaffID)

		switch tmp.ContentType {
		case model.MessageContentTypeText:
			tmp.RoomID = mc.RoomID
			mc.Notifier.Broadcast(tmp, mc, staff)
		case model.MessageContentTypeImage:
			tmp.RoomID = mc.RoomID
			mc.Notifier.Broadcast(tmp, mc, staff)
		case model.MessageContentTypeScore:
			mc.Notifier.MemberScored(mc, staff)
		}
	}
}

func (mc *MemberClient) SocketWrite() {
	defer func() {
		mc.Manager.unregister <- mc
	}()

	for {
		message, ok := <-mc.SendChan
		if !ok {
			if err := mc.Socket.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
				log.Error().Msgf("member write close message error: %s\n", err)
			}
			return
		}

		if err := mc.Socket.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Error().Msgf("member write message error: %s\n", err)
			return
		}
	}
}

// Reset 初始化從 sync.Pool 中拿出的 MemberClient
func (mc *MemberClient) Reset(clientInfo pkg.ClientInfo, manager *ClientManager) error {
	mc.ID = clientInfo.ID
	mc.Type = pkg.ClientTypeMember
	mc.Name = clientInfo.Name
	mc.RoomID = clientInfo.RoomID
	mc.StaffID = clientInfo.StaffID
	mc.Socket = clientInfo.Conn
	mc.SendChan = make(chan []byte)
	mc.Status = ClientStatusOpen
	mc.Manager = manager
	mc.dispatcher = manager.dispatcher
	mc.Notifier = manager.notifier
	mc.MsgSvc = manager.msgSvc
	mc.isSend = false
	mc.Sending = make(chan struct{})
	return nil
}
