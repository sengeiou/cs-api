package pkg

import (
	"cs-api/db/model"
	"cs-api/pkg/types"
	"github.com/gorilla/websocket"
)

type ClientType string

const (
	ClientTypeStaff  ClientType = "staff"
	ClientTypeMember ClientType = "member"
)

type ClientInfo struct {
	ID            int64                    `json:"id"`
	Type          ClientType               `json:"type"`
	Name          string                   `json:"name"`
	Username      string                   `json:"username"`       // staff only
	ServingStatus types.StaffServingStatus `json:"serving_status"` // staff only
	RoleID        int64                    `json:"role_id"`        // staff only
	Permissions   []string                 `json:"permissions"`    // staff only
	RoomID        int64                    `json:"room_id"`        // member only
	StaffID       int64                    `json:"staff_id"`       // member only
	Token         string                   `json:"token"`
	Conn          *websocket.Conn          `json:"conn,omitempty"`
}

type StaffEvent string

const (
	StaffEventClosed       StaffEvent = "Closed"
	StaffEventServing      StaffEvent = "Serving"
	StaffEventPending      StaffEvent = "Pending"
	StaffEventCloseRoom    StaffEvent = "CloseRoom"
	StaffEventAcceptRoom   StaffEvent = "AcceptRoom"
	StaffEventUpdateConfig StaffEvent = "UpdateConfig"
	StaffEventTransferRoom StaffEvent = "TransferRoom"
)

type StaffEventPayload struct {
	StaffID  *int64          `json:"staff_id,omitempty"`
	RoomID   *int64          `json:"room_id,omitempty"`
	CsConfig *types.CsConfig `json:"cs_config,omitempty"`
}

type StaffEventInfo struct {
	Event   StaffEvent        `json:"event"`
	Payload StaffEventPayload `json:"payload"`
}

type FastReplyCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type FastReplyGroupItem struct {
	Category FastReplyCategory                   `json:"category"`
	Items    []model.GetAllAvailableFastReplyRow `json:"items"`
}

type DailyTagReportColumn struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}
