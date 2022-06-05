package merchant

import (
	"context"
	"cs-api/db/model"
	"cs-api/dist/mock"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"reflect"
	"testing"
)

func Test_service_ListMerchant(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx          context.Context
		params       model.ListMerchantParams
		filterParams types.FilterMerchantParams
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantMerchants []model.ListMerchantRow
		wantCount     int64
		wantErr       bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:          context.Background(),
				params:       model.ListMerchantParams{},
				filterParams: types.FilterMerchantParams{},
			},
			wantMerchants: make([]model.ListMerchantRow, 0),
			wantCount:     0,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotTags, gotCount, err := s.ListMerchant(tt.args.ctx, tt.args.params, tt.args.filterParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMerchant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTags, tt.wantMerchants) {
				t.Errorf("ListMerchant() gotTags = %v, want %v", gotTags, tt.wantMerchants)
			}
			if gotCount != tt.wantCount {
				t.Errorf("ListMerchant() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_service_GetMerchant(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx        context.Context
		merchantId int64
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantMerchant model.GetMerchantRow
		wantErr      bool
	}{
		{
			name:   "normal test",
			fields: fields{repo: mock.NewRepository(t)},
			args: args{
				ctx:        context.Background(),
				merchantId: 1,
			},
			wantMerchant: model.GetMerchantRow{},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			gotTag, err := s.GetMerchant(tt.args.ctx, tt.args.merchantId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMerchant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTag, tt.wantMerchant) {
				t.Errorf("GetMerchant() gotTag = %v, want %v", gotTag, tt.wantMerchant)
			}
		})
	}
}

func Test_service_CreateMerchant(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.CreateMerchantParams
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
				params: model.CreateMerchantParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.CreateMerchant(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("CreateMerchant() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateMerchant(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx    context.Context
		params model.UpdateMerchantParams
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
				params: model.UpdateMerchantParams{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.UpdateMerchant(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UpdateMerchant() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteMerchant(t *testing.T) {
	type fields struct {
		repo iface.IRepository
	}
	type args struct {
		ctx        context.Context
		merchantId int64
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
				ctx:        context.Background(),
				merchantId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.DeleteMerchant(tt.args.ctx, tt.args.merchantId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMerchant() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ListAvailableMerchant(t *testing.T) {
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
		want    []model.ListAvailableMerchantRow
		wantErr bool
	}{
		{
			name:    "normal test",
			fields:  fields{repo: mock.NewRepository(t)},
			args:    args{ctx: context.Background()},
			want:    make([]model.ListAvailableMerchantRow, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.ListAvailableMerchant(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAvailableMerchant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAvailableMerchant() got = %v, want %v", got, tt.want)
			}
		})
	}
}
