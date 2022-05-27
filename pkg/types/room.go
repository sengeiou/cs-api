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
