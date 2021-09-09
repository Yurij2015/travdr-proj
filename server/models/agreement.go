package models

import "gorm.io/gorm"

type Agreement struct {
	Id            uint   `json:"id"`
	AgreementName string `json:"agreement_name"`
	Description   string `json:"description"`
	Content       string `json:"content"`
}

func (user *Agreement) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Agreement{}).Count(&total)

	return total
}

func (user *Agreement) Take(db *gorm.DB, limit int, offset int) interface{} {
	var agreements []Agreement
	db.Offset(offset).Limit(limit).Find(&agreements)
	return agreements
}
