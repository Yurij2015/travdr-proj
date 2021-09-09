package controllers

import (
	"../database"
	"../middlewares"
	"../models"
	"encoding/csv"
	"github.com/gofiber/fiber"
	"os"
	"strconv"
)

func AllCustomers(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "customers"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.Customer{}, page))
}

func AllCustomerList(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "customers"); err != nil {
		return err
	}
	var customers []models.Customer
	database.DB.Find(&customers)
	return c.JSON(customers)
}


func ExportCustomers(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "customers"); err != nil {
		return err
	}
	filePath := "./csv/customers.csv"
	if err := CreateFile(filePath); err != nil {
		return err
	}
	return c.Download(filePath)
}

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var customers []models.Customer

	database.DB.Find(&customers)

	writer.Write([]string{
		"ID", "FullName", "Gender", "BirthDate", "BirthPlace",
	})
	for _, customer := range customers {
		data := []string{
			strconv.Itoa(int(customer.Id)),
			customer.FullName,
			customer.Gender,
			customer.BirthDate,
			customer.BirthPlace,
		}
		if err := writer.Write(data); err != nil {
			return err
		}
	}
	return nil
}

// CRUD

func CreateCustomer(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "customers"); err != nil {
		return err
	}
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return err
	}
	database.DB.Create(&customer)
	return c.JSON(customer)
}

func GetCustomer(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "customers"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	customer := models.Customer{
		Id: uint(id),
	}
	database.DB.Find(&customer)
	return c.JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "customers"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	customer := models.Customer{
		Id: uint(id),
	}
	if err := c.BodyParser(&customer); err != nil {
		return err
	}
	database.DB.Model(&customer).Updates(customer)
	return c.JSON(customer)
}

func DeleteCustomer(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "customers"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	customer := models.Customer{
		Id: uint(id),
	}
	database.DB.Delete(&customer)
	return nil
}
