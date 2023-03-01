package user

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Fatal(err)
	}
	DNS := os.Getenv("DATABASE")

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot Connect to Database")
	}
	DB.AutoMigrate(&User{})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.Find(&user, id)
	return c.JSON(&user)
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	DB.Find(&users)
	return c.JSON(&users)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&user)
	return c.JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User Not available")
	}
	DB.Delete(&user)
	return c.SendString("User Deleted!")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(User)
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User Not available")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&user)
	return c.JSON(&user)
}
