package types

type RoomStatus int8

const (
	RoomStatusAll RoomStatus = iota
	RoomStatusPending
	RoomStatusServing
	RoomStatusClosed
)

type FilterRoomParams struct {
	RoomID  *int64
	StaffID *int64
	Status  *RoomStatus
}

type FilterStaffRoomParams struct {
	StaffID *int64
}

type Room struct {
	ID         int64      `json:"id"`
	Status     RoomStatus `json:"status"`
	CreatedAt  JSONTime   `json:"created_at"`
	ClosedAt   *JSONTime  `json:"closed_at"`
	StaffName  string     `json:"staff_name"`
	MemberName string     `json:"member_name"`
	TagName    string     `json:"tag_name"`
}
