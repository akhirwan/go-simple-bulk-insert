package config

import (
	"fmt"
	"go-simple-bulk-insert/infrastructure/database"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Env      string
	App      App
	Database database.MySQLConfig
}

type App struct {
	Name    string
	Version string
	Port    int
}

func LoadENVConfig() (config EnvironmentConfig, err error) {
	// Load .env file
	if err = godotenv.Load(); err != nil {
		err = fmt.Errorf("Read .env is failed: %s", err)
		return
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		err = fmt.Errorf("error when convert string to int: %s", err)
		return
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		err = fmt.Errorf("error when convert string to int: %s", err)
		return
	}

	config = EnvironmentConfig{
		Env: "DEV",
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
			Port:    port,
		},
		Database: database.MySQLConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		},
	}

	return
}
