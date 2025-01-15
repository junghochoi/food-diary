package models

type Entry struct {
	ID         string
  UserId     string
	Title      string
	Foods      []string
	FoodDesc   string
	Rating     uint8
	RatingDesc string
}
