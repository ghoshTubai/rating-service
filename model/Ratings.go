package model

type UserRatingData struct {
	ID      string   `json:"id"`
	Ratings []Rating `json:"rating"`
}

type Rating struct {
	MovieName string `json:"movieName"`
	Rating    int    `json:"rating"`
}
