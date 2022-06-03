package chat

import (
	"context"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"sync"
)

type StaffDispatcher struct {
	staffs    map[int64]*StaffClient // staffID map to conn
	maxMember int64
	lock      *sync.RWMutex
}

func (sd *StaffDispatcher) getStaff(staffId int64) *StaffClient {
	if staffId == 0 {
		return nil
	}

	return sd.staffs[staffId]
}

func (sd *StaffDispatcher) register(sc *StaffClient) {
	sd.staffs[sc.ID] = sc
}

func (sd *StaffDispatcher) unregister(sc *StaffClient) {
	delete(sd.staffs, sc.ID)
}

func (sd *StaffDispatcher) dispatch(staffId int64) *StaffClient {
	// 這一步是為了如果用戶不小心斷線重新連回來時，找回原本的客服
	staff := sd.getStaff(staffId)
	if staff != nil {
		return staff
	}

	// find available candidate
	var candidate []*StaffClient

	for _, client := range sd.staffs {
		if client.ServingStatus == types.StaffServingStatusServing && len(client.Rooms) < int(sd.maxMember) {
			candidate = append(candidate, client)
		}
	}

	staffCount := len(candidate)
	if staffCount == 0 {
		return nil
	}

	var tmp *StaffClient
	if staffCount == 1 {
		tmp = candidate[0]
	} else {
		// find staff has min memberCount
		tmp = candidate[0]
		for i := 1; i < staffCount; i++ {
			if len(candidate[i].Rooms) < len(tmp.Rooms) {
				tmp = candidate[i]
			}
		}
	}

	return tmp
}

func (sd *StaffDispatcher) setMaxMember(maxMember int64) {
	sd.maxMember = maxMember
}

func (sd *StaffDispatcher) assignRoom(staffId int64, roomId int64) {
	if staff, ok := sd.staffs[staffId]; ok {
		for _, v := range staff.Rooms {
			if v == roomId {
				return
			}
		}
		staff.Rooms = append(staff.Rooms, roomId)
	}
}

func (sd *StaffDispatcher) removeRoom(staffId int64, roomId int64) {
	if staff, ok := sd.staffs[staffId]; ok {
		idx := 0
		for k, v := range staff.Rooms {
			if v == roomId {
				idx = k
				break
			}
		}
		staff.Rooms = append(staff.Rooms[:idx], staff.Rooms[idx+1:]...)
	}
}

func (sd *StaffDispatcher) setServingStatus(staffId int64, servingStatus types.StaffServingStatus) {
	if staff, ok := sd.staffs[staffId]; ok {
		staff.ServingStatus = servingStatus
	}
}

func NewStaffDispatcher(csConfigSvc iface.ICsConfigService) *StaffDispatcher {
	config, err := csConfigSvc.GetCsConfig(context.Background())
	if err != nil {
		return nil
	}

	return &StaffDispatcher{
		staffs:    make(map[int64]*StaffClient),
		maxMember: config.MaxMember,
		lock:      &sync.RWMutex{},
	}
}
