package tag

import (
	iface "cs-api/pkg/interface"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewService,
		NewTagHandler,
	),
	fx.Invoke(
		InitTransport,
	),
)

type ITagHandler interface {
	ListTag(c *gin.Context)
	CreateTag(c *gin.Context)
	GetTag(c *gin.Context)
	UpdateTag(c *gin.Context)
	DeleteTag(c *gin.Context)
}

func NewTagHandler(authSvc iface.IAuthService, tagSvc iface.ITagService) ITagHandler {
	h := NewHttpHandler(authSvc, tagSvc)
	h = NewInstrumenting(h)
	return h
}
