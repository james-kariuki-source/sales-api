package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/james-kariuki-source/sales-api/connection"
	"github.com/james-kariuki-source/sales-api/routes"
)

func main() {
	db.Connect()

	app := fiber.New()
	app.Use(app)

	routes.Setup(app)

	app.Listen(":8080")
}
