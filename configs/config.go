package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DBConfig
}

type DBConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return &Config{
		Db: DBConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}
