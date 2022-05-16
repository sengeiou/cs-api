package message

import (
	"context"
	"cs-api/pkg"
	"cs-api/pkg/model"
	iface "github.com/golang/go-util/interface"
	mockTool "github.com/golang/go-util/mock"
	"reflect"
	"testing"
)

func Test_service_CreateMessage(t *testing.T) {
	type fields struct {
		repo iface.IMongoRepository
	}
	type args struct {
		ctx     context.Context
		message model.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mockTool.NewMongoRepository(t)},
			args: args{
				ctx:     context.Background(),
				message: model.Message{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.CreateMessage(tt.args.ctx, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("CreateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListRoomMessage(t *testing.T) {
	type fields struct {
		repo iface.IMongoRepository
	}
	type args struct {
		ctx        context.Context
		roomId     int64
		clientType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Message
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mockTool.NewMongoRepository(t)},
			args: args{
				ctx:        context.Background(),
				roomId:     1,
				clientType: "Text",
			},
			want:    make([]model.Message, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.ListRoomMessage(tt.args.ctx, tt.args.roomId, tt.args.clientType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRoomMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListRoomMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_ListMessage(t *testing.T) {
	type fields struct {
		repo iface.IMongoRepository
	}
	type args struct {
		ctx    context.Context
		params pkg.ListMessageParams
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantMessages []model.Message
		wantCount    int64
		wantErr      bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mockTool.NewMongoRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: pkg.ListMessageParams{},
			},
			wantMessages: make([]model.Message, 0),
			wantCount:    0,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotMessages, gotCount, err := s.ListMessage(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMessages, tt.wantMessages) {
				t.Errorf("ListMessage() gotMessages = %v, want %v", gotMessages, tt.wantMessages)
			}
			if gotCount != tt.wantCount {
				t.Errorf("ListMessage() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
