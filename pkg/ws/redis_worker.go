package ws

import (
	"context"
	"cs-api/pkg"
	"cs-api/pkg/types"
	"encoding/json"
	"github.com/golang/go-util/helper"
	iface "github.com/golang/go-util/interface"
	"github.com/rs/zerolog/log"
	"reflect"
)

func InitRedisSubscriber(redis iface.IRedis, worker *RedisWorker) {
	subscriber := redis.Subscribe(context.Background(), "event:staff")

	var info pkg.StaffEventInfo

	go func() {
		defer helper.Recover(context.Background())
		for {
			msg, err := subscriber.ReceiveMessage(context.Background())
			if err != nil {
				log.Error().Msgf("error: %s\n", err.Error())
				return
			}

			if err = json.Unmarshal([]byte(msg.Payload), &info); err != nil {
				log.Error().Msgf("error: %s\n", err.Error())
				return
			}

			worker.Handle(info)
		}
	}()
}

type RedisWorker struct {
	handler *EventHandler
	methods map[pkg.StaffEvent]reflect.Method
}

func NewRedisWorker(manager *ClientManager, dispatcher *StaffDispatcher, notifier *Notifier) *RedisWorker {
	worker := &RedisWorker{
		handler: &EventHandler{manager: manager, dispatcher: dispatcher, notifier: notifier},
		methods: make(map[pkg.StaffEvent]reflect.Method),
	}
	worker.registerMethods()
	return worker
}

func (r *RedisWorker) Handle(info pkg.StaffEventInfo) {
	f := r.methods[info.Event].Func
	f.Call([]reflect.Value{reflect.ValueOf(r.handler), reflect.ValueOf(info)})
}

func (r *RedisWorker) registerMethods() {
	t := reflect.TypeOf(r.handler)
	for i := 0; i < t.NumMethod(); i++ {
		name := t.Method(i).Name
		r.methods[pkg.StaffEvent(name)] = t.Method(i)
	}
}

type EventHandler struct {
	manager    *ClientManager
	dispatcher *StaffDispatcher
	notifier   *Notifier
}

func (r *EventHandler) CloseRoom(eventInfo pkg.StaffEventInfo) {
	r.manager.CloseRoom(*eventInfo.Payload.RoomID)
	r.dispatcher.removeRoom(*eventInfo.Payload.StaffID, *eventInfo.Payload.RoomID)
}

func (r *EventHandler) AcceptRoom(eventInfo pkg.StaffEventInfo) {
	member := r.manager.GetMember(*eventInfo.Payload.RoomID)
	// 如果用戶還沒下線
	if member != nil {
		member.StaffID = *eventInfo.Payload.StaffID
		r.notifier.RoomAccepted(member)
	}
}

func (r *EventHandler) UpdateConfig(eventInfo pkg.StaffEventInfo) {
	config := *eventInfo.Payload.CsConfig
	r.manager.csConfig = config
	r.dispatcher.setMaxMember(config.MaxMember)
}

func (r *EventHandler) TransferRoom(eventInfo pkg.StaffEventInfo) {
	member := r.manager.GetMember(*eventInfo.Payload.RoomID)
	member.StaffID = *eventInfo.Payload.StaffID
	staff := r.dispatcher.getStaff(member.StaffID)
	r.notifier.MemberJoin(member, staff)
	r.notifier.RoomTransferred(member)
}

func (r *EventHandler) Closed(eventInfo pkg.StaffEventInfo) {
	r.dispatcher.setServingStatus(*eventInfo.Payload.StaffID, types.StaffServingStatusClosed)
}

func (r *EventHandler) Serving(eventInfo pkg.StaffEventInfo) {
	r.dispatcher.setServingStatus(*eventInfo.Payload.StaffID, types.StaffServingStatusServing)
}

func (r *EventHandler) Pending(eventInfo pkg.StaffEventInfo) {
	r.dispatcher.setServingStatus(*eventInfo.Payload.StaffID, types.StaffServingStatusPending)
}
