package controller

import (
	"strconv"
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
		"success": true,
		"message": "Cashier has been added successfully",
		"data":    cashier,
	})

}

func RetrieveCashier(c *fiber.Ctx) error {
	var cashier []models.Cashier
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Here is a list of the cashiers",
			"data":    cashier,
		},
	)
}

func UpdateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Find(&cashier, "id=?", cashierId)

	if cashier.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": " cashier not found",
			},
		)
	}

	var UpdateCashier models.Cashier
	err := c.BodyParser(&UpdateCashier)
	if err != nil {
		return err
	}

	if UpdateCashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
		})
	}

	cashier.Name = UpdateCashier.Name
	db.DB.Save(&cashier)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier info updated successfully.",
	})
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Where("id=?", cashierId).First(&cashier)
	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	db.DB.Where("id=?", cashierId).Delete(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier successfully deleted.",
	})
}

func CashierDetails(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Select("id, name, created_at, updated_at").Where("id = ?", cashierId).First(&cashier)

	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["createdAt"] = cashier.CreatedAt
	cashierData["updatedAt"] = cashier.UpdatedAt

	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "cashier not found",
				"error":   map[string]interface{}{},
			},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier found",
			"data":    cashierData,
		},
	)
}
