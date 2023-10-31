package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func genRandomChar(charset string, length int) string {
	ranStr := make([]byte, length)
	for i := 0; i < length; i++ {
		ranStr[i] = charset[rand.Intn(length)]
	}
	return string(ranStr)
}
func redirectURL(c *fiber.Ctx) error {
	url := c.Params("url")
	if len(url) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "the length of url is less than 6 which is not possible",
		})
	}
	conn, err := connectDB()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "internal server error: no database connection",
		})
	}
	res, err := conn.Get(context.Background(), url).Result()
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		log.Println(res)
		return c.Status(fiber.StatusTemporaryRedirect).Redirect(res)
	}
}

func generateURL(c *fiber.Ctx) error {
	charset, exists := os.LookupEnv("CHARSET")
	if !exists {
		charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	}

	length := 6
	var requestBody requestData
	if err := c.BodyParser(&requestBody); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "invalid body",
		})
	}
	conn, err := connectDB()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "internal server error: no database connection",
		})
	}
	u, err := url.Parse(requestBody.Url)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "invalid url",
		})
	}
	u.Scheme = "https"
	requestBody.Url = u.String()
	result, err := conn.Get(context.Background(), requestBody.Url).Result()

	if err == nil {
		return c.Status(fiber.StatusOK).SendString("<a href=\"" + c.BaseURL() + "/" + result + "\">" + c.BaseURL() + "/" + result + "</a>")
	}

	var ranStr string
	for {
		ranStr = genRandomChar(charset, length)
		exists, err := conn.Exists(context.Background(), ranStr).Result()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   "internal server error: no database connection",
			})
		}
		if exists != 1 {
			break
		}
	}
	fmt.Println("the given url is ", requestBody.Url)
	_, err = conn.Set(context.Background(), ranStr, requestBody.Url, time.Duration(-1)).Result()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "internal server error: no database connection",
		})
	}
	_, err = conn.Set(context.Background(), requestBody.Url, ranStr, time.Duration(-1)).Result()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "internal server error: no database connection",
		})
	}

	return c.Status(fiber.StatusOK).SendString("<a href=\"" + c.BaseURL() + "/" + ranStr + "\">" + c.BaseURL() + "/" + ranStr + "</a>")
}

func indexPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
