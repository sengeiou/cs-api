package types

type StaffServingStatus int8

const (
	StaffServingStatusClosed StaffServingStatus = iota + 1
	StaffServingStatusServing
	StaffServingStatusPending
)

type FilterStaffParams struct {
	Name          *string
	Status        *Status
	ServingStatus *StaffServingStatus
}
