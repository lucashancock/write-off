package models

import (
	"gorm.io/gorm"
)

type Journal struct {
	gorm.Model
	Title   string
	Content string
}
