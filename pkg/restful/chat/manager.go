package chat

import (
	"context"
	"cs-api/db/model"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"sync"
	"time"
)

type ClientManager struct {
	lock          *sync.RWMutex // 需要加鎖防止等待用戶超出上限
	memberClients *sync.Map     // roomId maps member client
	dispatcher    *StaffDispatcher
	notifier      *Notifier
	roomSvc       iface.IRoomService
	msgSvc        iface.IMessageService
	memberSvc     iface.IMemberService
	csConfig      types.CsConfig
	memberPool    *sync.Pool
	staffPool     *sync.Pool
}

func (w *ClientManager) Register(clientInfo pkg.ClientInfo) {
	switch clientInfo.Type {
	case pkg.ClientTypeStaff:
		staff := w.staffPool.Get().(*StaffClient)
		if err := staff.Reset(clientInfo, w); err != nil {
			log.Error().Msgf("reset staff error: %s", err)
			return
		}
		w.dispatcher.register(staff)
		go staff.SocketRead()
		go staff.SocketWrite()
	case pkg.ClientTypeMember:
		// prevent duplicate member client connection
		if old := w.GetMember(clientInfo.RoomID); old != nil {
			if err := w.Unregister(old); err != nil {
				log.Error().Msgf("unregister member error: %s", err)
				return
			}
		}

		member := w.memberPool.Get().(*MemberClient)
		if err := member.Reset(clientInfo, w); err != nil {
			log.Error().Msgf("reset member error: %s", err)
			return
		}

		if err := w.memberSvc.UpdateOnlineStatus(context.Background(), model.UpdateOnlineStatusParams{
			OnlineStatus: types.MemberOnlineStatusOnline,
			ID:           member.ID,
		}); err != nil {
			log.Error().Msgf("update member online status error: %s", err)
			return
		}

		w.JoinRoom(member)

		// 檢查用戶是否超過閒置時間
		go checkPendingTimeout(member)
		go member.SocketRead()
		go member.SocketWrite()

		// 指派一位客服給用戶
		staff := w.dispatcher.dispatch(member.StaffID)
		if staff == nil {
			w.notifier.NoStaff(member)
		} else {
			member.StaffID = staff.ID
			if err := w.roomSvc.AcceptRoom(context.Background(), staff.ID, member.RoomID); err != nil {
				log.Error().Msgf("accept room error: %s", err.Error())
				return
			}
			w.dispatcher.assignRoom(staff.ID, member.RoomID)
			w.notifier.Greeting(w.csConfig.GreetingText, member, staff)
			w.notifier.MemberJoin(member, staff)
		}
	}
}

func (w *ClientManager) Unregister(client Client) error {
	if client.GetStatus() == ClientStatusClosed {
		return nil
	}
	if client.GetType() == pkg.ClientTypeStaff {
		staff := client.(*StaffClient)
		staff.Status = ClientStatusClosed
		close(staff.SendChan)
		if err := staff.Socket.Close(); err != nil {
			log.Error().Msgf("close socket error: %s\n", err)
			return err
		}
		w.dispatcher.unregister(staff)
		w.staffPool.Put(staff)
	} else if client.GetType() == pkg.ClientTypeMember {
		member := client.(*MemberClient)
		member.Status = ClientStatusClosed
		close(member.SendChan)
		close(member.Sending)
		if err := member.Socket.Close(); err != nil {
			log.Error().Msgf("handle unregister error: %s", err)
			return err
		}
		if err := w.memberSvc.UpdateOnlineStatus(context.Background(), model.UpdateOnlineStatusParams{
			OnlineStatus: types.MemberOnlineStatusOffline,
			ID:           member.ID,
		}); err != nil {
			log.Error().Msgf("update member online status error: %s", err)
			return err
		}
		w.memberClients.Delete(member.RoomID)
		w.memberPool.Put(member)
	}

	return nil
}

func (w *ClientManager) GetMember(roomId int64) *MemberClient {
	if member, ok := w.memberClients.Load(roomId); ok {
		return member.(*MemberClient)
	}
	return nil
}

func (w *ClientManager) JoinRoom(member *MemberClient) {
	w.memberClients.Store(member.RoomID, member)
}

func (w *ClientManager) CloseRoom(roomId int64) {
	if client, ok := w.memberClients.Load(roomId); ok {
		w.notifier.RoomClosed(client.(*MemberClient))
		time.Sleep(100 * time.Millisecond)
		if err := w.Unregister(client.(*MemberClient)); err != nil {
			log.Error().Msgf("unregister staff client error: %s", err)
			return
		}
	}
}

type ClientManagerParams struct {
	fx.In

	RoomSvc     iface.IRoomService
	MsgSvc      iface.IMessageService
	MemberSvc   iface.IMemberService
	CsConfigSvc iface.ICsConfigService
	Dispatcher  *StaffDispatcher
	Notifier    *Notifier
}

func NewClientManager(p ClientManagerParams) *ClientManager {
	config, err := p.CsConfigSvc.GetCsConfig(context.Background())
	if err != nil {
		return nil
	}

	return &ClientManager{
		lock:          &sync.RWMutex{}, // 需要加鎖防止等待用戶超出上限
		memberClients: &sync.Map{},     // roomId maps member client
		dispatcher:    p.Dispatcher,
		notifier:      p.Notifier,
		roomSvc:       p.RoomSvc,
		msgSvc:        p.MsgSvc,
		memberSvc:     p.MemberSvc,
		csConfig:      config,
		memberPool:    &sync.Pool{New: func() any { return &MemberClient{} }},
		staffPool:     &sync.Pool{New: func() any { return &StaffClient{} }},
	}
}

func checkPendingTimeout(member *MemberClient) {
	manager := member.Manager
	d := time.Duration(manager.csConfig.MemberPendingExpire) * time.Minute
	select {
	case <-time.After(d):
		manager.notifier.RoomClosed(member)
		time.Sleep(1 * time.Second)
		if err := manager.Unregister(member); err != nil {
			log.Error().Msgf("unregister staff client error: %s", err)
			return
		}
	case <-member.Sending:
		return
	}
}
