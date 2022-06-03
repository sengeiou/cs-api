package member

import (
	"cs-api/db/model"
	iface "cs-api/pkg/interface"
	"cs-api/pkg/types"
	"github.com/AndySu1021/go-util/errors"
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

type UpdateMemberStatusParams struct {
	Status types.MemberStatus `json:"status" binding:"required,oneof=1 2"`
}

func (h *handler) UpdateMemberStatus(c *gin.Context) {
	var (
		err           error
		requestParams UpdateMemberStatusParams
		ctx           = c.Request.Context()
	)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	if err = c.ShouldBindJSON(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	params := model.UpdateMemberStatusParams{
		Status: requestParams.Status,
		ID:     id,
	}

	if err = h.memberSvc.UpdateMemberStatus(ctx, params); err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.Success(c)
}

func NewHandler(authSvc iface.IAuthService, memberSvc iface.IMemberService) *handler {
	return &handler{
		authSvc:   authSvc,
		memberSvc: memberSvc,
	}
}
