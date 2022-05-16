package converter

import (
	"cs-api/db/model"
	"cs-api/pkg/types"
)

var RoomStatusMapping = map[RoomStatus]types.RoomStatus{
	RoomStatusAll:     0,
	RoomStatusPending: types.RoomStatusPending,
	RoomStatusServing: types.RoomStatusServing,
	RoomStatusClosed:  types.RoomStatusClosed,
}

var RoomStatusDtoMapping = map[types.RoomStatus]RoomStatus{
	types.RoomStatusPending: RoomStatusPending,
	types.RoomStatusServing: RoomStatusServing,
	types.RoomStatusClosed:  RoomStatusClosed,
}

func (input ListRoomInput) ConvertToModel(pagination PaginationInput) (model.ListRoomParams, types.FilterRoomParams) {
	params := model.ListRoomParams{
		Limit:  int32(pagination.PageSize),
		Offset: 0,
	}

	filterParams := types.FilterRoomParams{
		RoomID:  nil,
		StaffID: nil,
		Status:  nil,
	}

	if input.RoomID != 0 {
		filterParams.RoomID = &input.RoomID
	}

	if input.StaffID != 0 {
		filterParams.StaffID = &input.StaffID
	}

	if input.Status != RoomStatusAll {
		status := RoomStatusMapping[input.Status]
		filterParams.Status = &status
	}

	if pagination.Page > 0 {
		params.Offset = int32((pagination.Page - 1) * pagination.PageSize)
	}

	return params, filterParams
}

func (resp *ListRoomResp) FromModel(rooms []model.ListRoomRow, pagination PaginationInput, total int64) {
	resp.Rooms = make([]*Room, 0)

	for _, room := range rooms {
		tmp := Room{
			ID:         room.ID,
			MemberName: room.MemberName,
			StaffName:  room.StaffName,
			TagName:    room.TagName,
			StartTime:  room.CreatedAt.Format("2006-01-02 15:04:05"),
			EndTime:    "",
			Status:     RoomStatusDtoMapping[room.Status],
		}
		if room.ClosedAt.Valid != false {
			tmp.EndTime = room.ClosedAt.Time.Format("2006-01-02 15:04:05")
		}

		resp.Rooms = append(resp.Rooms, &tmp)
	}

	resp.Pagination = &Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    total,
	}
}

func (input ListStaffRoomInput) ConvertToModel(staffId int64, pagination PaginationInput) (model.ListStaffRoomParams, types.FilterStaffRoomParams) {
	params := model.ListStaffRoomParams{
		Status: RoomStatusMapping[input.Status],
		Limit:  int32(pagination.PageSize),
		Offset: 0,
	}

	filterParams := types.FilterStaffRoomParams{
		StaffID: nil,
	}

	if input.Status == RoomStatusServing {
		filterParams.StaffID = &staffId
	}

	if pagination.Page > 0 {
		params.Offset = int32((pagination.Page - 1) * pagination.PageSize)
	}

	return params, filterParams
}

func (resp *ListStaffRoomResp) FromModel(rooms []model.ListStaffRoomRow, pagination PaginationInput, total int64) {
	resp.Rooms = make([]*Room, 0)

	for _, room := range rooms {
		tmp := Room{
			ID:         room.ID,
			MemberName: room.MemberName,
			Status:     RoomStatusDtoMapping[room.Status],
		}

		resp.Rooms = append(resp.Rooms, &tmp)
	}

	resp.Pagination = &Pagination{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    total,
	}
}
