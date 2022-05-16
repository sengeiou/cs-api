package ws

import (
	"context"
	"cs-api/pkg"
	iface2 "cs-api/pkg/interface"
	"cs-api/pkg/model"
	"cs-api/pkg/types"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type StaffClient struct {
	ID            int64
	Type          pkg.WsClientType
	Name          string
	Socket        *websocket.Conn
	SendChan      chan []byte // 欲傳送的data
	Rooms         []int64     // 當前服務尚未關閉的房間ID
	Status        ClientStatus
	ServingStatus types.StaffServingStatus
	Manager       *ClientManager
	Notifier      *Notifier
	MsgSvc        iface2.IMessageService
}

func (sc *StaffClient) GetID() int64 {
	return sc.ID
}

func (sc *StaffClient) GetName() string {
	return sc.Name
}

func (sc *StaffClient) GetType() pkg.WsClientType {
	return sc.Type
}

func (sc *StaffClient) GetMessageType() model.MessageType {
	return model.MessageTypeStaff
}

func (sc *StaffClient) GetStatus() ClientStatus {
	return sc.Status
}

func (sc *StaffClient) GetSendChan() chan []byte {
	return sc.SendChan
}

func (sc *StaffClient) SocketRead() {
	defer func() {
		sc.Manager.unregister <- sc
	}()

	for {
		_, message, err := sc.Socket.ReadMessage()
		if err != nil {
			log.Error().Msgf("staff client ws error: %s\n", err)
			break
		}

		var tmp ClientMessage
		if err = json.Unmarshal(message, &tmp); err != nil {
			log.Error().Msgf("error: %s\n", err)
			break
		}

		member := sc.Manager.GetMember(tmp.RoomID)

		switch tmp.ContentType {
		case model.MessageContentTypeText:
			sc.Notifier.Broadcast(tmp, sc, member)
		case model.MessageContentTypeTyping:
			sc.Notifier.Typing(sc.Name, member)
		case model.MessageContentTypeScore:
			sc.Notifier.SendScore(tmp.RoomID, sc, member)
		}
	}
}

func (sc *StaffClient) SocketWrite() {
	defer func() {
		sc.Manager.unregister <- sc
	}()

	for {
		message, ok := <-sc.SendChan
		if !ok {
			if err := sc.Socket.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
				log.Error().Msgf("staff client write close message error: %s\n", err)
			}
			return
		}

		if err := sc.Socket.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Error().Msgf("staff client write message error: %s\n", err)
			return
		}
	}
}

// Reset 初始化從 sync.Pool 中拿出的 StaffClient
func (sc *StaffClient) Reset(clientInfo pkg.ClientInfo, manager *ClientManager) error {
	rooms, err := manager.roomSvc.GetStaffRooms(context.Background(), clientInfo.ID)
	if err != nil {
		log.Error().Msgf("get staff room id list error: %s", err.Error())
		return err
	}

	sc.ID = clientInfo.ID
	sc.Type = pkg.WsClientTypeStaff
	sc.Name = clientInfo.Name
	sc.Socket = clientInfo.Conn
	sc.SendChan = make(chan []byte)
	sc.Rooms = rooms
	sc.Status = ClientStatusOpen
	sc.ServingStatus = clientInfo.ServingStatus
	sc.Manager = manager
	sc.Notifier = manager.notifier
	sc.MsgSvc = manager.msgSvc

	return nil
}
