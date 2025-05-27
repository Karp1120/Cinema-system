package handlers

import (
	"encoding/json"
	"net/http"

	"cinema.local/db"
)

// 1. Репертуар по кинотеатру
func RepertoireByCinema(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("cinema")
	if name == "" {
		http.Error(w, "Не указано название кинотеатра", http.StatusBadRequest)
		return
	}

	rows, err := db.GetDB().Query(`
		SELECT f.title, s.date, s.time
		FROM sessions s
		JOIN cinemas c ON s.cinema_id = c.id
		JOIN films f ON s.film_id = f.id
		WHERE c.name ILIKE $1
		ORDER BY s.date, s.time
	`, name)
	if err != nil {
		http.Error(w, "Ошибка запроса", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Item struct {
		FilmTitle string `json:"film_title"`
		Date      string `json:"date"`
		Time      string `json:"time"`
	}

	var result []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.FilmTitle, &i.Date, &i.Time); err != nil {
			http.Error(w, "Ошибка чтения", http.StatusInternalServerError)
			return
		}
		result = append(result, i)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// 2. Кинотеатры по жанру
func CinemasByGenre(w http.ResponseWriter, r *http.Request) {
	genre := r.URL.Query().Get("genre")
	if genre == "" {
		http.Error(w, "Не указан жанр", http.StatusBadRequest)
		return
	}

	rows, err := db.GetDB().Query(`
		SELECT DISTINCT c.name, c.address
		FROM sessions s
		JOIN cinemas c ON s.cinema_id = c.id
		JOIN films f ON s.film_id = f.id
		WHERE f.genre ILIKE $1
	`, genre)
	if err != nil {
		http.Error(w, "Ошибка запроса", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Cinema struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	var result []Cinema
	for rows.Next() {
		var c Cinema
		if err := rows.Scan(&c.Name, &c.Address); err != nil {
			http.Error(w, "Ошибка чтения", http.StatusInternalServerError)
			return
		}
		result = append(result, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// 3. Свободные места и цена
func SessionInfo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Не указан ID сеанса", http.StatusBadRequest)
		return
	}

	row := db.GetDB().QueryRow(`SELECT price, available_seats FROM sessions WHERE id = $1`, id)

	type Info struct {
		Price     int `json:"price"`
		FreeSeats int `json:"free_seats"`
	}

	var info Info
	if err := row.Scan(&info.Price, &info.FreeSeats); err != nil {
		http.Error(w, "Сеанс не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// 4. Фильмы по режиссёру
func FilmsByDirector(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Не указано имя режиссёра", http.StatusBadRequest)
		return
	}

	rows, err := db.GetDB().Query(`
		SELECT title, genre FROM films WHERE director ILIKE $1
	`, name)
	if err != nil {
		http.Error(w, "Ошибка запроса", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Film struct {
		Title string `json:"title"`
		Genre string `json:"genre"`
	}

	var result []Film
	for rows.Next() {
		var f Film
		if err := rows.Scan(&f.Title, &f.Genre); err != nil {
			http.Error(w, "Ошибка чтения", http.StatusInternalServerError)
			return
		}
		result = append(result, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
