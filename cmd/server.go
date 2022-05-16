package cmd

import (
	"context"
	"cs-api/config"
	"cs-api/db/model"
	"cs-api/pkg/graph/resolver"
	"cs-api/pkg/lua"
	"cs-api/pkg/repository"
	"cs-api/pkg/restful"
	"cs-api/pkg/service"
	"cs-api/pkg/storage"
	"cs-api/pkg/ws"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/golang/go-util/db"
	"github.com/golang/go-util/gin"
	"github.com/golang/go-util/helper"
	"github.com/golang/go-util/mongo"
	"github.com/golang/go-util/redis"
	zlog "github.com/golang/go-util/zerolog"
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
		service.Module,
		resolver.Module,
		restful.Module,
		ws.Module,
	)

	exitCode := 0
	if err := app.Start(context.Background()); err != nil {
		log.Err(err).Msg("app start err")
		os.Exit(exitCode)
		return
	}

	//go func() {
	//	ticker := time.NewTicker(1 * time.Second)
	//	for {
	//		select {
	//		case <-ticker.C:
	//			var m runtime.MemStats
	//			runtime.ReadMemStats(&m)
	//			fmt.Printf("%v：memory = %.3f MB, GC Times = %v, Goroutine = %d\n", "Result", float64(m.Alloc)/1024/1024, m.NumGC, runtime.NumGoroutine())
	//		}
	//	}
	//}()

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
