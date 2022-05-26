package tag

import (
	"cs-api/pkg"
	iface "cs-api/pkg/interface"
	"github.com/gin-gonic/gin"
)

func InitTransport(engine *gin.Engine, authSvc iface.IAuthService, h ITagHandler) {
	routes := engine.Group("/api", authSvc.SetClientInfo(pkg.ClientTypeStaff))

	routes.GET("/tags", authSvc.CheckPermission("Tag.List"), h.ListTag)
	routes.POST("/tag", authSvc.CheckPermission("Tag.Create"), h.CreateTag)
	routes.GET("/tag/:id", authSvc.CheckPermission("Tag.GET"), h.GetTag)
	routes.PUT("/tag/:id", authSvc.CheckPermission("Tag.Update"), h.UpdateTag)
	routes.DELETE("/tag/:id", authSvc.CheckPermission("Tag.Delete"), h.DeleteTag)
}
