package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	db "github.com/james-kariuki-source/sales-api/connection"
	"github.com/james-kariuki-source/sales-api/routes"
)

func main() {
	db.Connect()

	app := fiber.New()
	app.Use(cors.New())

	routes.Setup(app)

	app.Listen(":8080")
}
