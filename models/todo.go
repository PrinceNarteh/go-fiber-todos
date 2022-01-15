package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Id        int    `gorm:"primaryKey"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
