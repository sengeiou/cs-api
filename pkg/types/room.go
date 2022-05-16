package types

type RoomStatus int8

const (
	RoomStatusPending RoomStatus = 1
	RoomStatusServing RoomStatus = 2
	RoomStatusClosed  RoomStatus = 3
)

type FilterRoomParams struct {
	RoomID  *int64
	StaffID *int64
	Status  *RoomStatus
}

type FilterStaffRoomParams struct {
	StaffID *int64
}
