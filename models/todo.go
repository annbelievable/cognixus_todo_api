package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title    string `gorm:"index;not null" json:"title"`
	Complete bool   `gorm:"index;not null;default:false" json:"complete"`
}
