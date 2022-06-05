package config

import (
	"github.com/AndySu1021/go-util/db"
	"github.com/AndySu1021/go-util/gin"
	"github.com/AndySu1021/go-util/log"
	"github.com/AndySu1021/go-util/redis"
	"github.com/AndySu1021/go-util/storage"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	Salt string `mapstructure:"salt"`
}

// AppConfig APP設定
type AppConfig struct {
	fx.Out

	App      *Config         `mapstructure:"app"`
	Storage  *storage.Config `mapstructure:"storage"`
	Http     *gin.Config     `mapstructure:"http"`
	Log      *log.Config     `mapstructure:"log"`
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

	return config, nil
}
