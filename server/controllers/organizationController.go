package controllers

import (
	"../database"
	"../middlewares"
	"../models"
	"github.com/gofiber/fiber"
	"strconv"
)

func AllOrganizations(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "organizations"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.Organization{}, page))
}
