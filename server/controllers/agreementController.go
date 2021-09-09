package controllers

import (
	"../database"
	"../middlewares"
	"../models"
	"github.com/gofiber/fiber"
	"strconv"
)

func AllAgreements(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agreements"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.Agreement{}, page))
}

func CreateAgreement(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agreements"); err != nil {
		return err
	}
	var agreement models.Agreement
	if err := c.BodyParser(&agreement); err != nil {
		return err
	}
	database.DB.Create(&agreement)
	return c.JSON(agreement)
}

func GetAgreement(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agreements"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	agreement := models.Agreement{
		Id: uint(id),
	}
	database.DB.Find(&agreement)
	return c.JSON(agreement)
}

func UpdateAgreement(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agreements"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	agreement := models.Agreement{
		Id: uint(id),
	}
	if err := c.BodyParser(&agreement); err != nil {
		return err
	}
	database.DB.Model(&agreement).Updates(agreement)
	return c.JSON(agreement)
}

func DeleteAgreement(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agreements"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	agent := models.Agreement{
		Id: uint(id),
	}
	database.DB.Delete(&agent)
	return nil
}