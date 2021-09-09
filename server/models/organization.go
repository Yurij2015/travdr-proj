package models

import "gorm.io/gorm"

type Organization struct {
	Id   uint   `json:"id"`
	OrgName string `json:"org_name"`
	OrgPlace string `json:"org_place"`
	OrgAddress string `json:"org_address"`
}

func (user *Organization) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Organization{}).Count(&total)

	return total
}

func (user *Organization) Take(db *gorm.DB, limit int, offset int) interface{} {
	var organizations []Organization
	db.Offset(offset).Limit(limit).Find(&organizations)
	return organizations
}