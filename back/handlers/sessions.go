package handlers

import (
	"encoding/json"
	"net/http"

	"cinema.local/db"
	"cinema.local/models"
)

// GET /api/sessions
func GetSessions(w http.ResponseWriter, r *http.Request) {
	rows, err := db.GetDB().Query(`SELECT id, cinema_id, film_id, date, time, price, available_seats FROM sessions`)
	if err != nil {
		http.Error(w, "Ошибка при получении списка сеансов", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sessions []models.Session
	for rows.Next() {
		var s models.Session
		err := rows.Scan(&s.ID, &s.CinemaID, &s.FilmID, &s.Date, &s.Time, &s.Price, &s.AvailableSeats)
		if err != nil {
			http.Error(w, "Ошибка при чтении строки", http.StatusInternalServerError)
			return
		}
		sessions = append(sessions, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// POST /api/sessions
func CreateSession(w http.ResponseWriter, r *http.Request) {
	var s models.Session
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	_, err := db.GetDB().Exec(
		`INSERT INTO sessions (cinema_id, film_id, date, time, price, available_seats) VALUES ($1, $2, $3, $4, $5, $6)`,
		s.CinemaID, s.FilmID, s.Date, s.Time, s.Price, s.AvailableSeats,
	)
	if err != nil {
		http.Error(w, "Ошибка при добавлении сеанса", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
