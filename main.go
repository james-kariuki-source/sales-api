package main

import (
	"fmt"
	"os"

	models "github.com/james-kariuki-source/sales-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	godotenv.Load()
	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic("Database connection failed!")
	}

	DB = db

	fmt.Println("The databse connection was successful!")

	AutoMigrate(db)

}

func AutoMigrate(connection *gorm.DB){
	connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Discount{},
		&models.Product{},
		&models.Order{},
	)
}
