package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})
	app.Get("/hash/:input", func(c *fiber.Ctx) error {
		input := c.Params("input")
		var hash [16]byte
		for i := 0; i < 100000; i++ {
			hash = md5.Sum([]byte(input))
		}
		return c.SendString(hex.EncodeToString(hash[:]))
	})
	app.Listen(":80")
}
