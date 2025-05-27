package handlers

import (
	"encoding/json"
	"net/http"

	"cinema.local/db"
	"cinema.local/models"
)

// GET /api/cinemas
func GetCinemas(w http.ResponseWriter, r *http.Request) {
	rows, err := db.GetDB().Query(`SELECT id, name, address, category, halls, seats, status FROM cinemas`)
	if err != nil {
		http.Error(w, "Ошибка при получении списка кинотеатров", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cinemas []models.Cinema
	for rows.Next() {
		var c models.Cinema
		err := rows.Scan(&c.ID, &c.Name, &c.Address, &c.Category, &c.Halls, &c.Seats, &c.Status)
		if err != nil {
			http.Error(w, "Ошибка при чтении строки", http.StatusInternalServerError)
			return
		}
		cinemas = append(cinemas, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cinemas)
}

// POST /api/cinemas
func CreateCinema(w http.ResponseWriter, r *http.Request) {
	var c models.Cinema
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	_, err := db.GetDB().Exec(
		`INSERT INTO cinemas (name, address, category, halls, seats, status) VALUES ($1, $2, $3, $4, $5, $6)`,
		c.Name, c.Address, c.Category, c.Halls, c.Seats, c.Status,
	)
	if err != nil {
		http.Error(w, "Ошибка при добавлении кинотеатра", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
