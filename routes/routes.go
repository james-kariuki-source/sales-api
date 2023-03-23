package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/james-kariuki-source/sales-api/controller"
)

func Setup(app *fiber.App){
	// app.Post("cashiers/:cashierId/login", controller.login)
	// app.Get("cashiers/:cashierId/login", controller.logout)
	// app.Post("cashiers/cashierId/passcode", controller.passcode)

	//Cashier routes
	app.Post("/cashiers", controller.CreateCashier)
	app.Get("/cashiers", controller.RetrieveCashier)
	app.Put("/cashiers/:cashierId", controller.UpdateCashier)
	app.Delete("/cashiers/:cashierId", controller.DeleteCashier)
	app.Get("/cashiers/:cashierId", controller.CashierDetails)
}