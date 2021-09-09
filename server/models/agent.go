package models

import "gorm.io/gorm"

type Agent struct {
	Id   uint   `json:"id"`
	FullName string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	OrganizationId uint `json:"organization_id"`
	Organization Organization `json:"organization" gorm:"foreignKey:OrganizationId"`
	Image     string `json:"image"`
}

func (user *Agent) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Agent{}).Count(&total)

	return total
}

func (user *Agent) Take(db *gorm.DB, limit int, offset int) interface{} {
	var agents []Agent
	db.Preload("Organization").Offset(offset).Limit(limit).Find(&agents)
	return agents
}