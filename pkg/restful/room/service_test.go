package room

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	mockTool "github.com/AndySu1021/go-util/mock"
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
)

func Test_service_CreateRoom(t *testing.T) {
	type fields struct {
		redis     ifaceTool.IRedis
		lua       iface.ILusScript
		memberSvc iface.IMemberService
		repo      iface.IRepository
	}
	type args struct {
		ctx      context.Context
		deviceId string
		name     string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantRoom   model.Room
		wantMember model.Member
		wantErr    bool
	}{
		{
			name: "normal test - room exists",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:      context.Background(),
				deviceId: "deviceId",
				name:     "name",
			},
			wantRoom:   model.Room{ID: 1, MemberID: 1},
			wantMember: model.Member{ID: 1},
			wantErr:    false,
		},
		{
			name: "normal test - room doesn't exist",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:      context.Background(),
				deviceId: "deviceId",
				name:     "n",
			},
			wantRoom:   model.Room{ID: 1, MemberID: 2},
			wantMember: model.Member{ID: 2},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:     tt.fields.redis,
				lua:       tt.fields.lua,
				memberSvc: tt.fields.memberSvc,
				repo:      tt.fields.repo,
			}
			gotRoom, gotMember, err := s.CreateRoom(tt.args.ctx, tt.args.deviceId, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotRoom.ID, tt.wantRoom.ID)
			assert.Equal(t, gotRoom.StaffID, int64(0))
			assert.Equal(t, gotRoom.MemberID, tt.wantMember.ID)
			if !reflect.DeepEqual(gotMember, tt.wantMember) {
				t.Errorf("CreateRoom() gotMember = %v, want %v", gotMember, tt.wantMember)
			}
		})
	}
}

func Test_service_AcceptRoom(t *testing.T) {
	type fields struct {
		redis     ifaceTool.IRedis
		lua       iface.ILusScript
		memberSvc iface.IMemberService
		repo      iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
		roomId  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal test - pending room",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
				roomId:  1,
			},
			wantErr: false,
		},
		{
			name: "normal test - closed room",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
				roomId:  4,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:     tt.fields.redis,
				lua:       tt.fields.lua,
				memberSvc: tt.fields.memberSvc,
				repo:      tt.fields.repo,
			}
			if err := s.AcceptRoom(tt.args.ctx, tt.args.staffId, tt.args.roomId); (err != nil) != tt.wantErr {
				t.Errorf("AcceptRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_CloseRoom(t *testing.T) {
	type fields struct {
		redis     ifaceTool.IRedis
		lua       iface.ILusScript
		memberSvc iface.IMemberService
		repo      iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
		roomId  int64
		tagId   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal test - serving room",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
				roomId:  2,
				tagId:   1,
			},
			wantErr: false,
		},
		{
			name: "normal test - pending room",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
				roomId:  1,
				tagId:   1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:     tt.fields.redis,
				lua:       tt.fields.lua,
				memberSvc: tt.fields.memberSvc,
				repo:      tt.fields.repo,
			}
			if err := s.CloseRoom(tt.args.ctx, tt.args.staffId, tt.args.roomId, tt.args.tagId); (err != nil) != tt.wantErr {
				t.Errorf("CloseRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateRoomScore(t *testing.T) {
	type fields struct {
		redis     ifaceTool.IRedis
		lua       iface.ILusScript
		memberSvc iface.IMemberService
		repo      iface.IRepository
	}
	type args struct {
		ctx    context.Context
		roomId int64
		score  int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:    context.Background(),
				roomId: 1,
				score:  5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:     tt.fields.redis,
				lua:       tt.fields.lua,
				memberSvc: tt.fields.memberSvc,
				repo:      tt.fields.repo,
			}
			if err := s.UpdateRoomScore(tt.args.ctx, tt.args.roomId, tt.args.score); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRoomScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListRoom(t *testing.T) {
	type fields struct {
		redis     ifaceTool.IRedis
		lua       iface.ILusScript
		memberSvc iface.IMemberService
		repo      iface.IRepository
	}
	type args struct {
		ctx          context.Context
		params       model.ListRoomParams
		filterParams types.FilterRoomParams
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantRooms []types.RoomList
		wantCount int64
		wantErr   bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:          context.Background(),
				params:       model.ListRoomParams{},
				filterParams: types.FilterRoomParams{},
			},
			wantRooms: make([]types.RoomList, 0),
			wantCount: 0,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:     tt.fields.redis,
				lua:       tt.fields.lua,
				memberSvc: tt.fields.memberSvc,
				repo:      tt.fields.repo,
			}
			gotRooms, gotCount, err := s.ListRoom(tt.args.ctx, tt.args.params, tt.args.filterParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRooms, tt.wantRooms) {
				t.Errorf("ListRoom() gotRooms = %v, want %v", gotRooms, tt.wantRooms)
			}
			if gotCount != tt.wantCount {
				t.Errorf("ListRoom() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_service_ListStaffRoom(t *testing.T) {
	type fields struct {
		redis     ifaceTool.IRedis
		lua       iface.ILusScript
		memberSvc iface.IMemberService
		repo      iface.IRepository
	}
	type args struct {
		ctx          context.Context
		params       model.ListStaffRoomParams
		filterParams types.FilterStaffRoomParams
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantRooms []model.ListStaffRoomRow
		wantCount int64
		wantErr   bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:          context.Background(),
				params:       model.ListStaffRoomParams{},
				filterParams: types.FilterStaffRoomParams{},
			},
			wantRooms: make([]model.ListStaffRoomRow, 0),
			wantCount: 0,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:     tt.fields.redis,
				lua:       tt.fields.lua,
				memberSvc: tt.fields.memberSvc,
				repo:      tt.fields.repo,
			}
			gotRooms, gotCount, err := s.ListStaffRoom(tt.args.ctx, tt.args.params, tt.args.filterParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListStaffRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRooms, tt.wantRooms) {
				t.Errorf("ListStaffRoom() gotRooms = %v, want %v", gotRooms, tt.wantRooms)
			}
			if gotCount != tt.wantCount {
				t.Errorf("ListStaffRoom() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_service_GetStaffRooms(t *testing.T) {
	type fields struct {
		redis     ifaceTool.IRedis
		lua       iface.ILusScript
		memberSvc iface.IMemberService
		repo      iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int64
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
			},
			want:    make([]int64, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:     tt.fields.redis,
				lua:       tt.fields.lua,
				memberSvc: tt.fields.memberSvc,
				repo:      tt.fields.repo,
			}
			got, err := s.GetStaffRooms(tt.args.ctx, tt.args.staffId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaffRooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaffRooms() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_TransferRoom(t *testing.T) {
	type fields struct {
		redis     ifaceTool.IRedis
		lua       iface.ILusScript
		memberSvc iface.IMemberService
		repo      iface.IRepository
	}
	type args struct {
		ctx       context.Context
		staffId   int64
		roomId    int64
		toStaffId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:       context.Background(),
				staffId:   1,
				roomId:    2,
				toStaffId: 2,
			},
			wantErr: false,
		},
		{
			name: "staff not available",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:       context.Background(),
				staffId:   1,
				roomId:    1,
				toStaffId: 1,
			},
			wantErr: true,
		},
		{
			name: "no need to transfer",
			fields: fields{
				redis:     mockTool.NewRedis(t),
				lua:       mock.NewLuaScript(t),
				memberSvc: mock.NewMemberService(t),
				repo:      mock.NewRepository(t),
			},
			args: args{
				ctx:       context.Background(),
				staffId:   1,
				roomId:    3,
				toStaffId: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:     tt.fields.redis,
				lua:       tt.fields.lua,
				memberSvc: tt.fields.memberSvc,
				repo:      tt.fields.repo,
			}
			if err := s.TransferRoom(tt.args.ctx, tt.args.staffId, tt.args.roomId, tt.args.toStaffId); (err != nil) != tt.wantErr {
				t.Errorf("TransferRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
