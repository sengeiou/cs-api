package types

type StaffServingStatus int8

const (
	StaffServingStatusAll StaffServingStatus = iota
	StaffServingStatusClosed
	StaffServingStatusServing
	StaffServingStatusPending
)

type FilterStaffParams struct {
	Name          *string
	Status        *Status
	ServingStatus *StaffServingStatus
}
