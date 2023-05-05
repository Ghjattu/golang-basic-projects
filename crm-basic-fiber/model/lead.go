package model

import (
	"crm-basic-fiber/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Lead{})
}

func GetLeads(c *fiber.Ctx) error {
	var leads []Lead
	err := db.Find(&leads).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(leads)
}

func GetLeadByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if id == 0 || err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var lead Lead
	err = db.Where("ID=?", id).First(&lead).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	return c.JSON(lead)
}

func CreateLead(c *fiber.Ctx) error {
	var lead Lead
	err := c.BodyParser(&lead)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = db.Create(&lead).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if id == 0 || err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var lead Lead
	err = db.Where("ID=?", id).Delete(&lead).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(lead)
}
