package config

import (
	"github.com/spf13/viper"
	"miniService/internal/db"
	"os"
)

type Config struct {
	Port     string
	Postgres db.PostgresConfig
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("default")

	return viper.ReadInConfig()
}

func New() Config {
	return Config{
		Port: viper.GetString("port"),
		Postgres: db.PostgresConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			//Password: "password",
			DBName:  viper.GetString("db.dbname"),
			SSLMode: viper.GetString("db.sslmode"),
		},
	}
}
