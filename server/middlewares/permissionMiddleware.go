package middlewares

import (
	"../database"
	"../models"
	"../util"
	"errors"
	"github.com/gofiber/fiber"
	"strconv"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")
	Id, err := util.ParseJwt(cookie)
	if err != nil {
		return err
	}
	userId, _ := strconv.Atoi(Id)
	user := models.User{
		Id: uint(userId),
	}
	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}
	database.DB.Preload("Permissions").Find(&role)
	if c.Method() == "GET" {
		for _, permissions := range role.Permissions {
			if permissions.Name == "view_"+page || permissions.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permissions := range role.Permissions {
			if permissions.Name == "edit_"+page {
				return nil
			}
		}
	}
	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")
}
