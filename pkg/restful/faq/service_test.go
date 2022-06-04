package faq

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"reflect"
	"testing"
)

func Test_service_ListFAQ(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx          context.Context
		params       model.ListFAQParams
		filterParams types.FilterFAQParams
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantFAQs  []model.ListFAQRow
		wantCount int64
		wantErr   bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:          context.Background(),
				params:       model.ListFAQParams{},
				filterParams: types.FilterFAQParams{},
			},
			wantFAQs:  make([]model.ListFAQRow, 0),
			wantCount: 0,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotFAQs, gotCount, err := s.ListFAQ(tt.args.ctx, tt.args.params, tt.args.filterParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListFAQ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFAQs, tt.wantFAQs) {
				t.Errorf("ListFAQ() gotFAQs = %v, want %v", gotFAQs, tt.wantFAQs)
			}
			if gotCount != tt.wantCount {
				t.Errorf("ListFAQ() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_service_GetFAQ(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx   context.Context
		faqId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.GetFAQRow
		wantErr bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:   context.Background(),
				faqId: 1,
			},
			want:    model.GetFAQRow{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetFAQ(tt.args.ctx, tt.args.faqId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFAQ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFAQ() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateFAQ(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.CreateFAQParams
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
				params: model.CreateFAQParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.CreateFAQ(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateFAQ() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateFAQ(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.UpdateFAQParams
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
				params: model.UpdateFAQParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.UpdateFAQ(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateFAQ() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteFAQ(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx   context.Context
		faqId int64
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
				ctx:   context.Background(),
				faqId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.DeleteFAQ(tt.args.ctx, tt.args.faqId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFAQ() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListAvailableFAQ(t *testing.T) {
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
		want    []model.ListAvailableFAQRow
		wantErr bool
	}{
		{
			name:    "normal test",
			fields:  fields{repo: mock.NewRepository(t)},
			args:    args{ctx: context.Background()},
			want:    make([]model.ListAvailableFAQRow, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.ListAvailableFAQ(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAvailableFAQ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAvailableFAQ() got = %v, want %v", got, tt.want)
			}
		})
	}
}
