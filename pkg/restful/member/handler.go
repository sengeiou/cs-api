package member

import (
	iface "cs-api/pkg/interface"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"strconv"
)

type handler struct {
	authSvc   iface.IAuthService
	memberSvc iface.IMemberService
}

func (h *handler) GetMemberStatus(c *gin.Context) {
	var (
		err error
		ctx = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	status, err := h.memberSvc.GetMemberStatus(ctx, id)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, gin.H{
		"status": status,
	})
}

func NewHandler(authSvc iface.IAuthService, memberSvc iface.IMemberService) *handler {
	return &handler{
		authSvc:   authSvc,
		memberSvc: memberSvc,
	}
}
