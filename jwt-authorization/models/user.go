package models

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null; unique"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email"`
}

var db *gorm.DB

func init() {
	d, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}
	db = d
	db.AutoMigrate(&User{})
}

func CreateUser(user *User) (*User, error) {
	err := db.Create(user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

func LoginCheck(username, password string) (uint, uint) {
	var user User

	if err := db.Where("username=?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, fiber.StatusNotFound
		}
		return 0, fiber.StatusInternalServerError
	}

	if user.Password != password {
		return 0, fiber.StatusNotFound
	}

	return user.ID, fiber.StatusOK
}
