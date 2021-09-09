package controllers

import (
	"../database"
	"../middlewares"
	"../models"
	"github.com/gofiber/fiber"
	"strconv"
)

func AllContracts(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "contracts"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.Contract{}, page))
}

func CreateContract(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "contracts"); err != nil {
		return err
	}
	var contract models.Contract
	if err := c.BodyParser(&contract); err != nil {
		return err
	}
	database.DB.Create(&contract)
	return c.JSON(contract)
}

func GetContract(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "contracts"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	contract := models.Contract{
		Id: uint(id),
	}
	database.DB.Preload("Organization").Preload("Agreement").Preload("Agent").Preload("Customer").Find(&contract)
	return c.JSON(contract)
}

func UpdateContract(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "contracts"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	contract := models.Contract{
		Id: uint(id),
	}
	if err := c.BodyParser(&contract); err != nil {
		return err
	}
	database.DB.Model(&contract).Updates(contract)
	return c.JSON(contract)
}

func DeleteContract(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "contracts"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	agent := models.Contract{
		Id: uint(id),
	}
	database.DB.Delete(&agent)
	return nil
}