package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	db "github.com/james-kariuki-source/sales-api/connection"
	"github.com/james-kariuki-source/sales-api/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateCashier(c *fiber.Ctx) error {

	//Collecting and checking the cashier input data
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "The cashier must have a name",
			})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "The cashier should have a passcode",
			})
	}

	//hashing the passcode
	passcode, _ := bcrypt.GenerateFromPassword([]byte(data["passcode"]), 14)

	//Saving the cashier data
	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  passcode,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	db.DB.Create(&cashier)

	return c.Status(200).JSON(fiber.Map{
		"success":true,
		"message": "Cashier has been added successfully",
		"data": cashier,
	})


}

func RetrieveCashier(c *fiber.Ctx) error {
	return nil
}

func UpdateCashier(c *fiber.Ctx) error {
	return nil
}

func DeleteCashier(c *fiber.Ctx) error {
	return nil
}

func CashierDetails(c *fiber.Ctx) error {
	return nil
}
