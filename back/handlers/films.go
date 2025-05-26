package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cinema.local/db"
	"cinema.local/models"

	"github.com/gorilla/mux"
)

// GET /api/films
func GetFilms(w http.ResponseWriter, r *http.Request) {
	rows, err := db.GetDB().Query(`SELECT id, title, director, genre, studio, operator, actors FROM films`)
	if err != nil {
		http.Error(w, "Ошибка при получении фильмов", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var films []models.Film
	for rows.Next() {
		var f models.Film
		err := rows.Scan(&f.ID, &f.Title, &f.Director, &f.Genre, &f.Studio, &f.Operator, &f.Actors)
		if err != nil {
			http.Error(w, "Ошибка чтения данных", http.StatusInternalServerError)
			return
		}
		films = append(films, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(films)
}

// POST /api/films
func CreateFilm(w http.ResponseWriter, r *http.Request) {
	var f models.Film
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	_, err := db.GetDB().Exec(
		`INSERT INTO films (title, director, genre, studio, operator, actors) VALUES ($1, $2, $3, $4, $5, $6)`,
		f.Title, f.Director, f.Genre, f.Studio, f.Operator, f.Actors,
	)
	if err != nil {
		http.Error(w, "Ошибка при добавлении фильма", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DELETE /api/films/{id}
func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	_, err = db.GetDB().Exec(`DELETE FROM films WHERE id = $1`, id)
	if err != nil {
		http.Error(w, "Ошибка при удалении фильма", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
