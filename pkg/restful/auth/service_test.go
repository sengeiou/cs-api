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
		want    pkg.ClientInfo
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
			want: pkg.ClientInfo{
				ID:            1,
				Type:          "staff",
				Name:          "user",
				Username:      "user",
				ServingStatus: 1,
				Permissions:   []string{"*"},
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
			assert.Equal(t, got.Permissions, tt.want.Permissions)
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
		staffInfo pkg.ClientInfo
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
				staffInfo: pkg.ClientInfo{},
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

func Test_service_GetClientInfo(t *testing.T) {
	var (
		clientInfo = pkg.ClientInfo{
			ID:            1,
			Type:          pkg.ClientTypeStaff,
			Name:          "user",
			Username:      "user",
			ServingStatus: 1,
			Token:         "token",
		}
		ctx = context.WithValue(context.Background(), "staff_info", clientInfo)
	)
	type fields struct {
		redis  ifaceTool.IRedis
		lua    iface.ILusScript
		repo   iface.IRepository
		config *config.Config
	}
	type args struct {
		ctx        context.Context
		clientType pkg.ClientType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    pkg.ClientInfo
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
				ctx:        ctx,
				clientType: pkg.ClientTypeStaff,
			},
			want:    clientInfo,
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
			args: args{
				ctx:        context.Background(),
				clientType: "staff",
			},
			want:    pkg.ClientInfo{},
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
			got, err := s.GetClientInfo(tt.args.ctx, tt.args.clientType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClientInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClientInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
