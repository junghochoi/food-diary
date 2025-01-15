package request

type EntryGetRequest struct {
	ID string `json:"id"`
}

type EntryCreateRequest struct {
	Title      string   `json:"title"`
	Foods      []string `json:"foods"`
	FoodDesc   string   `json:"foodDesc"`
	Rating     uint8    `json:"rating"`
	RatingDesc string   `json:"ratingDesc"`
}
