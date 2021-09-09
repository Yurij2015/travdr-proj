package models

import "gorm.io/gorm"

type Employee struct {
	Id        uint   `json:"id"`
	ShortName string `json:"short_name"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birth_date"`
	Image     string `json:"image"`
}

func (employee *Employee) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Employee{}).Count(&total)

	return total
}

func (employee *Employee) Take(db *gorm.DB, limit int, offset int) interface{} {
	var employees []Employee
	db.Offset(offset).Limit(limit).Find(&employees)
	return employees
}
