package configuration

import (
	"fmt"
	"os"
)

type Config struct {
	DbUser      string
	DbPassword  string
	DbName      string
	DbHost      string
	DbPort      string
	DatabaseURL string
}

func LoadConfig() Config {
	config := Config{
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
	}

	config.DatabaseURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	return config
}
