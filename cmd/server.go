package cmd

import (
	"context"
	"cs-api/config"
	"cs-api/db/model"
	"cs-api/pkg/lua"
	"cs-api/pkg/repository"
	"cs-api/pkg/restful"
	"database/sql"
	"errors"
	"github.com/AndySu1021/go-util/db"
	"github.com/AndySu1021/go-util/gin"
	"github.com/AndySu1021/go-util/helper"
	"github.com/AndySu1021/go-util/mongo"
	"github.com/AndySu1021/go-util/redis"
	"github.com/AndySu1021/go-util/storage"
	zlog "github.com/AndySu1021/go-util/zerolog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ServerCmd = &cobra.Command{
	Run: runServer,
	Use: "server",
}

func runServer(_ *cobra.Command, _ []string) {
	defer helper.Recover(context.Background())

	logger := log.Level(zerolog.InfoLevel)

	commonModule := fx.Options(
		fx.Provide(
			config.NewConfig,
			db.NewDatabase,
			db.NewMongo,
			mongo.New,
			gin.NewGin,
			lua.NewLua,
			NewDBTX,
			model.New,
			repository.NewRepository,
			storage.NewStorage,
			NewContext,
		),
		fx.Invoke(
			zlog.InitZeroLog,
			Migrate,
			Seed,
		),
	)

	app := fx.New(
		fx.Logger(&logger),
		commonModule,
		redis.Module,
		restful.Module,
	)

	exitCode := 0
	if err := app.Start(context.Background()); err != nil {
		log.Err(err).Msg("app start err")
		os.Exit(exitCode)
		return
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-stopChan
	log.Info().Msgf("main: shutting down server...")

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Err(err).Msg("app stop err")
	}

	os.Exit(exitCode)
}

func NewDBTX(db *sql.DB) model.DBTX {
	return db
}

func Migrate(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "mysql", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return err
}

func Seed(cfg *config.Config, q *model.Queries) (err error) {
	ctx := context.Background()
	if err = q.TagSeeder(ctx); err != nil {
		return
	}
	if err = q.RoleSeeder(ctx); err != nil {
		return
	}
	if err = q.StaffSeeder(ctx, helper.EncryptPassword("admin", cfg.Salt)); err != nil {
		return
	}
	if err = q.ConstantSeeder(ctx); err != nil {
		return
	}
	return
}

func NewContext() context.Context {
	return context.Background()
}
