package models

import "gorm.io/gorm"

type Customer struct {
	Id         uint   `json:"id"`
	FullName   string `json:"full_name"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birth_date"`
	BirthPlace string `json:"birth_place"`
	Image     string `json:"image"`
}

func (user *Customer) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Customer{}).Count(&total)

	return total
}

func (user *Customer) Take(db *gorm.DB, limit int, offset int) interface{} {
	var customers []Customer
	db.Offset(offset).Limit(limit).Find(&customers)
	return customers
}
