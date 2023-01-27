package models

import (
	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	Deadline string `json:"deadline"`
	Name     string `json:"name"`
}
