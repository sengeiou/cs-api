package member

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/pkg/interface"
	"github.com/magiconair/properties/assert"
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
