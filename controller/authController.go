package controller

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	db "github.com/james-kariuki-source/sales-api/connection"
	"github.com/james-kariuki-source/sales-api/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid post request",
		})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Passcode is required",
			"error":   map[string]interface{}{},
		})
	}

	var cashier models.Cashier
	db.DB.Where("id=?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
			"error":   map[string]interface{}{},
		})
	}

	if err := bcrypt.CompareHashAndPassword(cashier.Passcode, []byte(data["passcode"])); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid passcode",
			"error":   map[string]interface{}{},
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issers":   strconv.Itoa(int(cashier.Id)),
		"ExpireAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Token expired or is invalid",
		})
	}

	cashierData := make(map[string]interface{})
	cashierData["token"] = tokenString

	return c.Status(200).JSON(fiber.Map{
		"success": false,
		"message": "Successful login",
		"data":    cashierData,
	})
}

func passcode(c fiber.Ctx) {

}

func Logout(c fiber.Ctx) {

}
