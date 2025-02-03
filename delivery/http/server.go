package http

import (
	"go-simple-bulk-insert/delivery/container"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ServeHttp(container container.Container) *fiber.App {
	log.Println("Starting server...")
	handler := SetupHandler(container)

	app := fiber.New() // iniate fiber context

	RouterGroup(app, handler) // iniate router v1

	return app
}
