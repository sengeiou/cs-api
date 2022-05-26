package restful

import (
	"cs-api/pkg/restful/auth"
	"cs-api/pkg/restful/cs_config"
	"cs-api/pkg/restful/member"
	"cs-api/pkg/restful/notice"
	"cs-api/pkg/restful/remind"
	"cs-api/pkg/restful/report"
	"cs-api/pkg/restful/role"
	"cs-api/pkg/restful/tag"
	"cs-api/pkg/restful/tool"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	fx.Provide(
		tool.NewRequestInstrument,
	),
	fx.Invoke(
		InitCommonHandler,
	),
	auth.Module,
	tag.Module,
	role.Module,
	notice.Module,
	remind.Module,
	member.Module,
	cs_config.Module,
	report.Module,
)

func InitCommonHandler(engine *gin.Engine) {
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	engine.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })
}
