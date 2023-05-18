package controllers

import (
	"jwt-authorization/models"
	"jwt-authorization/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error."})
	}

	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username and Password can not be empty."})
	}

	if _, err := models.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error."})
	}
	return c.JSON(fiber.Map{"message": "Registered Successfully."})
}

func Login(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error."})
	}

	userID, statusCode := models.LoginCheck(user.Username, user.Password)
	if statusCode == fiber.StatusNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Username or password is incorrect."})
	}

	token, err := utils.GenerateToken(userID)
	if err != nil || statusCode == fiber.StatusInternalServerError {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error."})
	}
	return c.JSON(fiber.Map{"token": token})
}

func GetAllUsers(c *fiber.Ctx) error {
	users, err := models.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error."})
	}
	return c.JSON(users)
}
