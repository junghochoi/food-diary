package models

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Title             string
	Foods             []string `gorm:"type:json"` // JSON column for storing a list of strings
	FoodsDescription  string
	Rating            uint8
	RatingDescription string
}
