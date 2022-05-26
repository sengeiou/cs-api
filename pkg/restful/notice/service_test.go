package notice

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"reflect"
	"testing"
)

func Test_service_ListNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx          context.Context
		params       model.ListNoticeParams
		filterParams types.FilterNoticeParams
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantNotices []model.Notice
		wantCount   int64
		wantErr     bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:          context.Background(),
				params:       model.ListNoticeParams{},
				filterParams: types.FilterNoticeParams{},
			},
			wantNotices: make([]model.Notice, 0),
			wantCount:   0,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotNotices, gotCount, err := s.ListNotice(tt.args.ctx, tt.args.params, tt.args.filterParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListNotice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNotices, tt.wantNotices) {
				t.Errorf("ListNotice() gotNotices = %v, want %v", gotNotices, tt.wantNotices)
			}
			if gotCount != tt.wantCount {
				t.Errorf("ListNotice() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_service_GetNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx      context.Context
		noticeId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Notice
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				noticeId: 1,
			},
			want:    model.Notice{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetNotice(tt.args.ctx, tt.args.noticeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNotice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNotice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.CreateNoticeParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: model.CreateNoticeParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.CreateNotice(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateNotice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.UpdateNoticeParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:    context.Background(),
				params: model.UpdateNoticeParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.UpdateNotice(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateNotice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx      context.Context
		noticeId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:      context.Background(),
				noticeId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.DeleteNotice(tt.args.ctx, tt.args.noticeId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteNotice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetAvailableNotice(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Notice
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx: context.Background(),
			},
			want:    model.Notice{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetLatestNotice(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAvailableNotice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvailableNotice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
