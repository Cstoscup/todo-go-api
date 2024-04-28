package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Task      string `gorm:"not null;type:text" json:"task"`
	Completed bool   `gorm:"default:false" json:"completed"`
}
