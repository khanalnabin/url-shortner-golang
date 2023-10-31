package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
