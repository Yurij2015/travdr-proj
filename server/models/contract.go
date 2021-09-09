package models

import "gorm.io/gorm"

type Contract struct {
	Id   uint   `json:"id"`
	ContractNumber string `json:"contract_number"`
	StartTripDate string `json:"start_trip_date"`
	EndTripDate string `json:"end_trip_date"`
	AgreementId uint `json:"agreement_id"`
	Agreement Agreement `json:"agreement" gorm:"foreignKey:AgreementId"`
	OrganizationId uint `json:"organization_id"`
	Organization Organization `json:"organization" gorm:"foreignKey:OrganizationId"`
	AgentId uint `json:"agent_id"`
	Agent Agent `json:"agent" gorm:"foreignKey:AgentId"`
	CustomerId uint `json:"customer_id"`
	Customer Customer `json:"customer" gorm:"foreignKey:CustomerId"`
}

func (user *Contract) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Contract{}).Count(&total)

	return total
}

func (user *Contract) Take(db *gorm.DB, limit int, offset int) interface{} {
	var agents []Contract
	db.Preload("Organization").Preload("Agreement").Preload("Agent").Preload("Customer").Offset(offset).Limit(limit).Find(&agents)
	return agents
}