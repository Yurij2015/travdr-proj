package database

import (
	"../models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:root@/travelers_dream"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to db")
	}
	DB = database
	err = database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Employee{}, &models.Organization{}, &models.Agent{}, &models.Customer{}, &models.Contract{}, &models.Agreement{})
	if err != nil {
		return
	}
}
