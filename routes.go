package main

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *fiber.App) {
	app.Get("/:url", redirectURL)
	app.Post("/", generateURL)
	app.Get("/", indexPage)
}
