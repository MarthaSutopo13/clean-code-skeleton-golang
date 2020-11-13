package config

import (
	"basesvc_v2/internal/config/db"
	"basesvc_v2/internal/config/logging"
	"basesvc_v2/internal/config/server"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	Server   server.ServerList
	Database db.DatabaseList
	Logger   logging.LoggerConfig
}

var cfg config

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(dir + "/internal/config/server")
	viper.SetConfigType("yaml")
	viper.SetConfigName("rest.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load server config: %v", err))
	}

	viper.AddConfigPath(dir + "/server")
	viper.SetConfigName("client.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load server client config: %v", err))
	}

	viper.AddConfigPath(dir + "/internal/config/logging")
	viper.SetConfigType("yaml")
	viper.SetConfigName("logger.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load server config: %v", err))
	}

	viper.AddConfigPath(dir + "/internal/config/db")
	viper.SetConfigName("postgres.yml")
	err = viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot read database config: %v", err))
	}

	viper.Unmarshal(&cfg)
	viper.AutomaticEnv()

	if viper.GetString("CARD_GRPC_URL") != "" {
		cfg.Server.Rest.Tracking.URL = viper.GetString("CARD_GRPC_URL")
	}
}

func GetConfig() *config {
	return &cfg
}
