package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cs-api/pkg/graph/converter"
)

func (r *queryResolver) ListDailyTagReport(ctx context.Context, filter converter.ListDailyTagReportInput) (*converter.ListDailyTagReportResp, error) {
	_, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return nil, err
	}

	columns, result, err := r.reportSvc.DailyTagReport(ctx, filter.StartDate, filter.EndDate)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListDailyTagReportResp{}
	resp.FromMap(columns, result)

	return resp, nil
}

func (r *queryResolver) ListDailyGuestReport(ctx context.Context, filter converter.ListDailyGuestReportInput) (*converter.ListDailyGuestReportResp, error) {
	_, err := r.authSvc.GetStaffInfo(ctx)
	if err != nil {
		return nil, err
	}

	result, err := r.reportSvc.DailyGuestReport(ctx, filter.StartDate, filter.EndDate)
	if err != nil {
		return nil, err
	}

	resp := &converter.ListDailyGuestReportResp{}
	resp.FromMap(result)

	return resp, nil
}
