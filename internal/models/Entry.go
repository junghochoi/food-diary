package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Title             string
	Foods             pq.StringArray `gorm:"type:text[]"`
	FoodsDescription  string
	Rating            uint8
	RatingDescription string
}
