package report

import (
	iface "cs-api/pkg/interface"
	"github.com/AndySu1021/go-util/errors"
	ginTool "github.com/AndySu1021/go-util/gin"
	"github.com/gin-gonic/gin"
	"time"
)

type handler struct {
	authSvc   iface.IAuthService
	reportSvc iface.IReportService
}

type DailyTagReportParams struct {
	StartDate time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" binding:"required" time_format:"2006-01-02"`
}

func (h *handler) DailyTagReport(c *gin.Context) {
	var (
		err           error
		requestParams DailyTagReportParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	column, data, err := h.reportSvc.DailyTagReport(ctx, requestParams.StartDate, requestParams.EndDate)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, gin.H{
		"columns": column,
		"items":   data,
	})
}

type DailyGuestReportParams struct {
	StartDate time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" binding:"required" time_format:"2006-01-02"`
}

func (h *handler) DailyGuestReport(c *gin.Context) {
	var (
		err           error
		requestParams DailyGuestReportParams
		ctx           = c.Request.Context()
	)

	if err = c.ShouldBindQuery(&requestParams); err != nil {
		ginTool.Error(c, errors.ErrorValidation)
		return
	}

	data, err := h.reportSvc.DailyGuestReport(ctx, requestParams.StartDate, requestParams.EndDate)
	if err != nil {
		ginTool.Error(c, err)
		return
	}

	ginTool.SuccessWithData(c, data)
}

func NewHandler(authSvc iface.IAuthService, reportSvc iface.IReportService) *handler {
	return &handler{
		authSvc:   authSvc,
		reportSvc: reportSvc,
	}
}
