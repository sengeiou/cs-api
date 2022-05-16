package cron

import (
	"context"
	"cs-api/config"
	"cs-api/db/model"
	"cs-api/pkg/service"
	"database/sql"
	"github.com/golang/go-util/db"
	"github.com/golang/go-util/helper"
	zlog "github.com/golang/go-util/zerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"os"
	"time"
)

var ReportDailyTagCmd = &cobra.Command{
	Run: runReportDailyTag,
	Use: "ReportDailyTag",
}

var commonModule = fx.Options(
	fx.Provide(
		config.NewConfig,
		db.NewDatabase,
		NewDBTX,
		model.New,
	),
	fx.Invoke(
		zlog.InitZeroLog,
		execReportDailyTag,
	),
)

var dateFlag string

func init() {
	ReportDailyTagCmd.PersistentFlags().StringVar(&dateFlag, "date", "", "Run report with specific date")
}

func runReportDailyTag(_ *cobra.Command, _ []string) {
	defer helper.Recover(context.Background())

	logger := log.Level(zerolog.InfoLevel)
	fxOption := []fx.Option{
		fx.Logger(&logger),
	}

	fxOption = append(fxOption, commonModule, service.Module)

	app := fx.New(
		fxOption...,
	)

	exitCode := 0
	if err := app.Start(context.Background()); err != nil {
		log.Err(err).Msg("app start err")
		os.Exit(exitCode)
		return
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Err(err).Msg("app stop err")
	}

	os.Exit(exitCode)
}

func execReportDailyTag(query *model.Queries) {
	date, startTime, endTime, err := GetTimeRange(dateFlag1)
	if err != nil {
		log.Error().Msgf("get time range error: %s", err)
		return
	}

	ctx := context.Background()
	if err = query.DeleteReportDailyTag(ctx, *date); err != nil {
		log.Error().Msgf("delete report daily tag error: %s", err)
		return
	}

	items, err := query.CountClosedRoomByTag(ctx, model.CountClosedRoomByTagParams{
		ClosedAt:   sql.NullTime{Time: *startTime, Valid: true},
		ClosedAt_2: sql.NullTime{Time: *endTime, Valid: true},
	})
	if err != nil {
		log.Error().Msgf("count closed room by tag error: %s", err)
		return
	}
	for _, item := range items {
		if err = query.CreateReportDailyTag(ctx, model.CreateReportDailyTagParams{
			Date:      *date,
			TagID:     item.TagID,
			Count:     int32(item.Count),
			CreatedAt: time.Now().UTC(),
		}); err != nil {
			log.Error().Msgf("create report daily tag error: %s", err)
			return
		}
	}
}
