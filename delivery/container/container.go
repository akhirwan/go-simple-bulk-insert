package container

import (
	"go-simple-bulk-insert/config"
	"go-simple-bulk-insert/domain/counter/feature"
	"go-simple-bulk-insert/domain/counter/repository"
	"go-simple-bulk-insert/infrastructure/database"
	"log"
)

type Container struct {
	EnvConfig      config.EnvironmentConfig
	CounterFeature feature.CounterFeature
}

func SetupContainer(config config.EnvironmentConfig) Container {
	log.Println("Connecting database...")
	db, err := database.NewMySQLDBConnection(&config.Database)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Loading repository's...")
	counterRepository := repository.NewCounterRepository(db)

	log.Println("Loading feature's...")
	counterFeature := feature.NewCounterFeature(counterRepository)

	return Container{
		EnvConfig:      config,
		CounterFeature: counterFeature,
	}
}
