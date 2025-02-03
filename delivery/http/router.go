package http

import "github.com/gofiber/fiber/v2"

func RouterGroup(app *fiber.App, handler handler) {
	app.Post("/create", handler.counterHandler.CreateHandler)
}
