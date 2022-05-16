package mock

import (
	iface "cs-api/pkg/interface"
	"github.com/golang/mock/gomock"
	"testing"
)

func NewLuaScript(t *testing.T) iface.ILusScript {
	m := gomock.NewController(t)
	mock := NewMockILusScript(m)

	mock.EXPECT().SetToken(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().RemoveToken(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	return mock
}
