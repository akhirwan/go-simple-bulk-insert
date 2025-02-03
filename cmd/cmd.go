package cmd

import (
	"fmt"
	"go-simple-bulk-insert/config"
	"go-simple-bulk-insert/delivery/container"
	"go-simple-bulk-insert/delivery/http"
	"log"
)

func Execute() {
	log.Println("Loading config...")
	config, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	// start init container
	container := container.SetupContainer(config)

	// start http service
	http := http.ServeHttp(container)
	http.Listen(fmt.Sprintf(":%d", config.App.Port))
}
