package handlers

import (
	"encoding/json"
	"net/http"

	"cinema.local/db"
	"cinema.local/models"
)

// GET /api/sessions
func GetSessions(w http.ResponseWriter, r *http.Request) {
	rows, err := db.GetDB().Query(`SELECT id, cinema_id, film_id, date, time, price, available_seats, hall_number FROM sessions`)
	if err != nil {
		http.Error(w, "Ошибка при получении списка сеансов", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sessions []models.Session
	for rows.Next() {
		var s models.Session
		err := rows.Scan(&s.ID, &s.CinemaID, &s.FilmID, &s.Date, &s.Time, &s.Price, &s.AvailableSeats, &s.HallNumber)
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

	// Проверка: в этом зале, в эту дату и время, в этом кинотеатре уже есть сеанс?
	var exists bool
	err := db.GetDB().QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM sessions
			WHERE cinema_id = $1 AND hall_number = $2 AND date = $3 AND time = $4
		)
	`, s.CinemaID, s.HallNumber, s.Date, s.Time).Scan(&exists)

	if err != nil {
		http.Error(w, "Ошибка при проверке расписания", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "Сеанс в этом зале на эту дату и время уже существует", http.StatusConflict)
		return
	}

	_, err = db.GetDB().Exec(
		`INSERT INTO sessions (cinema_id, film_id, date, time, price, available_seats, hall_number) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		s.CinemaID, s.FilmID, s.Date, s.Time, s.Price, s.AvailableSeats, s.HallNumber,
	)
	if err != nil {
		http.Error(w, "Ошибка при добавлении сеанса", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
