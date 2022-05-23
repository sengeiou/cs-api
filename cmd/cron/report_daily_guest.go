package cron

import (
	"context"
	"cs-api/config"
	"cs-api/db/model"
	"cs-api/pkg/service"
	"database/sql"
	"errors"
	"github.com/AndySu1021/go-util/db"
	"github.com/AndySu1021/go-util/helper"
	zlog "github.com/AndySu1021/go-util/zerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"os"
	"time"
)

var ReportDailyGuestCmd = &cobra.Command{
	Run: runReportDailyGuest,
	Use: "ReportDailyGuest",
}

var commonModule1 = fx.Options(
	fx.Provide(
		config.NewConfig,
		db.NewDatabase,
		NewDBTX,
		model.New,
	),
	fx.Invoke(
		zlog.InitZeroLog,
		execReportDailyGuest,
	),
)

var dateFlag1 string

func init() {
	ReportDailyGuestCmd.PersistentFlags().StringVar(&dateFlag1, "date", "", "Run report with specific date")
}

func runReportDailyGuest(_ *cobra.Command, _ []string) {
	defer helper.Recover(context.Background())

	logger := log.Level(zerolog.InfoLevel)
	fxOption := []fx.Option{
		fx.Logger(&logger),
	}

	fxOption = append(fxOption, commonModule1, service.Module)

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

func execReportDailyGuest(query *model.Queries) {
	date, startTime, endTime, err := GetTimeRange(dateFlag1)
	if err != nil {
		log.Error().Msgf("get time range error: %s", err)
		return
	}

	ctx := context.Background()
	if err = query.DeleteReportDailyGuest(ctx, *date); err != nil {
		log.Error().Msgf("delete report daily guest error: %s", err)
		return
	}

	result, err := query.CountDailyRoomByMember(ctx, model.CountDailyRoomByMemberParams{
		CreatedAt:   *startTime,
		CreatedAt_2: *endTime,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Error().Msgf("count daily room by member error: %s", err)
		return
	}

	if err = query.CreateReportDailyGuest(ctx, model.CreateReportDailyGuestParams{
		Date:       *date,
		GuestCount: int32(result),
		CreatedAt:  time.Now().UTC(),
	}); err != nil {
		log.Error().Msgf("create report daily guest error: %s", err)
		return
	}
}
