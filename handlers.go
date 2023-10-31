package main

import "github.com/gofiber/fiber/v2"

func redirectURL(c *fiber.Ctx) error {
	return nil
}

func generateURL(c *fiber.Ctx) error {
	return nil
}

func indexPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
