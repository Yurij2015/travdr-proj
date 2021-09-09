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

func AllAgents(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agents"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.Agent{}, page))
}

func AllAgentsList(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agents"); err != nil {
		return err
	}
	var agents []models.Agent
	database.DB.Find(&agents)
	return c.JSON(agents)
}

func ExportAgents(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agents"); err != nil {
		return err
	}
	filePath := "./csv/agents.csv"
	if err := CreateAgentsFile(filePath); err != nil {
		return err
	}
	return c.Download(filePath)
}

func CreateAgentsFile(filePath string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var agents []models.Agent

	database.DB.Preload("Organization").Find(&agents)

	writer.Write([]string{
		"ID", "FullName", "PhoneNumber", "Organization", "Image",
	})
	for _, agent := range agents {
		data := []string{
			strconv.Itoa(int(agent.Id)),
			agent.FullName,
			agent.PhoneNumber,
			agent.Organization.OrgName,
			agent.Image,
		}
		if err := writer.Write(data); err != nil {
			return err
		}
	}
	return nil
}

// CRUD

func CreateAgent(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agents"); err != nil {
		return err
	}
	var agent models.Agent
	if err := c.BodyParser(&agent); err != nil {
		return err
	}
	database.DB.Create(&agent)
	return c.JSON(agent)
}

func GetAgent(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agents"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	agent := models.Agent{
		Id: uint(id),
	}
	database.DB.Find(&agent)
	return c.JSON(agent)
}

func UpdateAgent(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agents"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	agent := models.Agent{
		Id: uint(id),
	}
	if err := c.BodyParser(&agent); err != nil {
		return err
	}
	database.DB.Model(&agent).Updates(agent)
	return c.JSON(agent)
}

func DeleteAgent(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "agents"); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Params("id"))
	agent := models.Agent{
		Id: uint(id),
	}
	database.DB.Delete(&agent)
	return nil
}
