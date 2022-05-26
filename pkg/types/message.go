package types

type ListMessageParams struct {
	RoomID   int64
	StaffID  int64
	Content  string
	Page     int64
	PageSize int64
}
