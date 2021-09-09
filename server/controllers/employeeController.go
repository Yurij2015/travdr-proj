package controllers

import (
	"../database"
	"../middlewares"
	"../models"
	"github.com/gofiber/fiber"
	"strconv"
)

func AllEmployees(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "employees"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.Employee{}, page))
}

func CreateEmployee(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "employees"); err != nil {
		return err
	}
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return err
	}
	database.DB.Create(&employee)
	return c.JSON(employee)
}

func GetEmployee(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "employees"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	employee := models.Employee{
		Id: uint(id),
	}
	database.DB.Find(&employee)
	return c.JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "employees"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	employee := models.Employee{
		Id: uint(id),
	}
	if err := c.BodyParser(&employee); err != nil {
		return err
	}
	database.DB.Model(&employee).Updates(employee)
	return c.JSON(employee)
}

func DeleteEmployee(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "employees"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	employee := models.Employee{
		Id: uint(id),
	}
	//if err := c.BodyParser(&employee); err != nil {
	//	return err
	//}
	database.DB.Delete(&employee)
	return nil
}
