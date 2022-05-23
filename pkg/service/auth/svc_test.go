package auth

import (
	"context"
	"cs-api/config"
	"cs-api/dist/mock"
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	ifaceTool "github.com/AndySu1021/go-util/interface"
	mockTool "github.com/AndySu1021/go-util/mock"
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
)

func Test_service_Login(t *testing.T) {
	type fields struct {
		redis  ifaceTool.IRedis
		lua    iface.ILusScript
		repo   iface.IRepository
		config *config.Config
	}
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    pkg.StaffInfo
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis:  mockTool.NewRedis(t),
				lua:    mock.NewLuaScript(t),
				repo:   mock.NewRepository(t),
				config: &config.Config{Salt: "salt"},
			},
			args: args{
				ctx:      context.Background(),
				username: "user",
				password: "user",
			},
			want: pkg.StaffInfo{
				ID:            1,
				Type:          1,
				Name:          "user",
				Username:      "user",
				ServingStatus: 1,
				Token:         "token",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:  tt.fields.redis,
				lua:    tt.fields.lua,
				repo:   tt.fields.repo,
				config: tt.fields.config,
			}
			got, err := s.Login(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, got.ID, tt.want.ID)
			assert.Equal(t, got.Type, tt.want.Type)
			assert.Equal(t, got.Name, tt.want.Name)
			assert.Equal(t, got.Username, tt.want.Username)
			assert.Equal(t, got.ServingStatus, tt.want.ServingStatus)
			assert.Matches(t, got.Token, "^[0-9a-z]{32}$")
		})
	}
}

func Test_service_Logout(t *testing.T) {
	type fields struct {
		redis  ifaceTool.IRedis
		lua    iface.ILusScript
		repo   iface.IRepository
		config *config.Config
	}
	type args struct {
		ctx       context.Context
		staffInfo pkg.StaffInfo
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
				redis:  mockTool.NewRedis(t),
				lua:    mock.NewLuaScript(t),
				repo:   mock.NewRepository(t),
				config: &config.Config{Salt: "salt"},
			},
			args: args{
				ctx:       context.Background(),
				staffInfo: pkg.StaffInfo{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:  tt.fields.redis,
				lua:    tt.fields.lua,
				repo:   tt.fields.repo,
				config: tt.fields.config,
			}
			if err := s.Logout(tt.args.ctx, tt.args.staffInfo); (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetStaffInfo(t *testing.T) {
	var (
		staffInfo = pkg.StaffInfo{
			ID:            1,
			Type:          1,
			Name:          "user",
			Username:      "user",
			ServingStatus: 1,
			Token:         "token",
		}
		ctx = context.WithValue(context.Background(), "staff_info", staffInfo)
	)
	type fields struct {
		redis  ifaceTool.IRedis
		lua    iface.ILusScript
		repo   iface.IRepository
		config *config.Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    pkg.StaffInfo
		wantErr bool
	}{
		{
			name: "normal test",
			fields: fields{
				redis:  mockTool.NewRedis(t),
				lua:    mock.NewLuaScript(t),
				repo:   mock.NewRepository(t),
				config: &config.Config{Salt: "salt"},
			},
			args:    args{ctx: ctx},
			want:    staffInfo,
			wantErr: false,
		},
		{
			name: "authentication error",
			fields: fields{
				redis:  mockTool.NewRedis(t),
				lua:    mock.NewLuaScript(t),
				repo:   mock.NewRepository(t),
				config: &config.Config{Salt: "salt"},
			},
			args:    args{ctx: context.Background()},
			want:    pkg.StaffInfo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				redis:  tt.fields.redis,
				lua:    tt.fields.lua,
				repo:   tt.fields.repo,
				config: tt.fields.config,
			}
			got, err := s.GetStaffInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaffInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaffInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
