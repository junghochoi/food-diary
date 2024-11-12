package models

type Entry struct {
	ID         string
	Title      string
	Foods      []string
	FoodDesc   string
	Rating     uint8
	RatingDesc string
}
