package main

import(
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	app.Get("/testing", Hello)

	app.Listen(":8000")
}

func Hello(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"success" : true,
		"message" : "This is a text of the api",

	})
}