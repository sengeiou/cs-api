package member

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
)

func Test_service_GetOrCreateMember(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx      context.Context
		name     string
		deviceId string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantMember model.Member
		wantErr    bool
	}{
		{
			name:   "normal test - guest member exists",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				name:     "",
				deviceId: "deviceId",
			},
			wantMember: model.Member{
				ID:       1,
				Type:     2,
				DeviceID: "deviceId",
			},
			wantErr: false,
		},
		{
			name:   "normal test - guest member doesn't exist",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				name:     "",
				deviceId: "device",
			},
			wantMember: model.Member{
				ID:       1,
				Type:     2,
				DeviceID: "device",
			},
			wantErr: false,
		},
		{
			name:   "normal test - normal member exists",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				name:     "name",
				deviceId: "deviceId",
			},
			wantMember: model.Member{
				ID:       1,
				Type:     1,
				Name:     "name",
				DeviceID: "deviceId",
			},
			wantErr: false,
		},
		{
			name:   "normal test - normal member doesn't exist",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				name:     "n",
				deviceId: "deviceId",
			},
			wantMember: model.Member{
				ID:       1,
				Type:     1,
				Name:     "n",
				DeviceID: "deviceId",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotMember, err := s.GetOrCreateMember(tt.args.ctx, tt.args.name, tt.args.deviceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrCreateMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotMember.ID, tt.wantMember.ID)
			assert.Equal(t, gotMember.Type, tt.wantMember.Type)
			if gotMember.Name == "" {
				assert.Matches(t, gotMember.Name, "^Guest-([0-9a-z]{3})$")
			} else {
				assert.Matches(t, gotMember.Name, tt.wantMember.Name)
			}
			assert.Equal(t, gotMember.DeviceID, tt.wantMember.DeviceID)
		})
	}
}

func Test_service_GetOnlineStatus(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx      context.Context
		memberId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    types.MemberOnlineStatus
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				memberId: 1,
			},
			want:    types.MemberOnlineStatus(1),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetOnlineStatus(tt.args.ctx, tt.args.memberId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMemberStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMemberStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateOnlineStatus(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.UpdateOnlineStatusParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: model.UpdateOnlineStatusParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.UpdateOnlineStatus(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateMemberStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListMember(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx          context.Context
		params       model.ListMemberParams
		filterParams types.FilterMemberParams
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantMembers []model.Member
		wantCount   int64
		wantErr     bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:          context.Background(),
				params:       model.ListMemberParams{},
				filterParams: types.FilterMemberParams{},
			},
			wantMembers: make([]model.Member, 0),
			wantCount:   0,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotMembers, gotCount, err := s.ListMember(tt.args.ctx, tt.args.params, tt.args.filterParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMembers, tt.wantMembers) {
				t.Errorf("ListMember() gotMembers = %v, want %v", gotMembers, tt.wantMembers)
			}
			if gotCount != tt.wantCount {
				t.Errorf("ListMember() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
