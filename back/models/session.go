package models

type Session struct {
	ID             int    `json:"id"`
	CinemaID       int    `json:"cinema_id"`
	FilmID         int    `json:"film_id"`
	Date           string `json:"date"`
	Time           string `json:"time"`
	Price          int    `json:"price"`
	AvailableSeats int    `json:"available_seats"`
}
