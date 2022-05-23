package staff

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	mockTool "github.com/AndySu1021/go-util/mock"
	"reflect"
	"testing"
)

func Test_service_ListStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx          context.Context
		params       model.ListStaffParams
		filterParams types.FilterStaffParams
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStaffs []model.ListStaffRow
		wantCount  int64
		wantErr    bool
	}{
		{
			name: "normal test",
			fields: fields{
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:          context.Background(),
				params:       model.ListStaffParams{},
				filterParams: types.FilterStaffParams{},
			},
			wantStaffs: make([]model.ListStaffRow, 0),
			wantCount:  0,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			gotStaffs, gotCount, err := s.ListStaff(tt.args.ctx, tt.args.params, tt.args.filterParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStaffs, tt.wantStaffs) {
				t.Errorf("ListStaff() gotStaffs = %v, want %v", gotStaffs, tt.wantStaffs)
			}
			if gotCount != tt.wantCount {
				t.Errorf("ListStaff() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_service_GetStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantStaff model.GetStaffRow
		wantErr   bool
	}{
		{
			name: "normal test",
			fields: fields{
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
			},
			wantStaff: model.GetStaffRow{ID: 1, ServingStatus: 1},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			gotStaff, err := s.GetStaff(tt.args.ctx, tt.args.staffId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStaff, tt.wantStaff) {
				t.Errorf("GetStaff() gotStaff = %v, want %v", gotStaff, tt.wantStaff)
			}
		})
	}
}

func Test_service_CreateStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.CreateStaffParams
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
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:    context.Background(),
				params: model.CreateStaffParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			if err := s.CreateStaff(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params interface{}
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
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:    context.Background(),
				params: model.UpdateStaffParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			if err := s.UpdateStaff(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
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
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			if err := s.DeleteStaff(tt.args.ctx, tt.args.staffId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteStaff() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateStaffServingStatus(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx       context.Context
		staffInfo pkg.StaffInfo
		status    types.StaffServingStatus
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
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:       context.Background(),
				staffInfo: pkg.StaffInfo{},
				status:    types.StaffServingStatusServing,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			if err := s.UpdateStaffServingStatus(tt.args.ctx, tt.args.staffInfo, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStaffServingStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListAvailableStaff(t *testing.T) {
	type fields struct {
		redis ifaceTool.IRedis
		repo  iface.IRepository
	}
	type args struct {
		ctx     context.Context
		staffId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Staff
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				repo:  mock.NewRepository(t),
				redis: mockTool.NewRedis(t),
			},
			args: args{
				ctx:     context.Background(),
				staffId: 1,
			},
			want:    make([]model.Staff, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis: tt.fields.redis,
				repo:  tt.fields.repo,
			}
			got, err := s.ListAvailableStaff(tt.args.ctx, tt.args.staffId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAvailableStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAvailableStaff() got = %v, want %v", got, tt.want)
			}
		})
	}
}
