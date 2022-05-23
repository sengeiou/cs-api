package config

import (
	"cs-api/pkg/types"
	"github.com/AndySu1021/go-util/db"
	"github.com/AndySu1021/go-util/gin"
	"github.com/AndySu1021/go-util/redis"
	zlog "github.com/AndySu1021/go-util/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	Salt        string           `mapstructure:"salt"`
	DiskDriver  types.DiskDriver `mapstructure:"disk_driver"`
	DiskBaseUrl string           `mapstructure:"disk_base_url"`
}

type DiskConfig struct {
	Driver  types.DiskDriver `mapstructure:"driver"`
	BaseUrl string           `mapstructure:"base_url"`
}

// AppConfig APP設定
type AppConfig struct {
	fx.Out

	App      *Config         `mapstructure:"app"`
	Disk     *DiskConfig     `mapstructure:"disk"`
	Http     *gin.Config     `mapstructure:"http"`
	Log      *zlog.Config    `mapstructure:"log"`
	Database *db.Config      `mapstructure:"database"`
	Mongo    *db.MongoConfig `mapstructure:"mongo"`
	Redis    *redis.Config   `mapstructure:"redis"`
}

// NewConfig Initiate config
func NewConfig() (AppConfig, error) {
	viper.AutomaticEnv()
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	var config = AppConfig{}

	if err := viper.ReadInConfig(); err != nil {
		return AppConfig{}, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return AppConfig{}, err
	}

	log.Debug().Msgf("salt is %s", config.App.Salt)

	return config, nil
}
