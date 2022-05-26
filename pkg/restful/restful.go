package restful

import (
	"cs-api/pkg/restful/tag"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	fx.Invoke(
		InitMetrics,
	),
	tag.Module,
)

func InitMetrics(engine *gin.Engine) {
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	engine.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })
}
