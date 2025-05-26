package main

import (
	"log"
	"net/http"

	gorillahandlers "github.com/gorilla/handlers"

	"github.com/gorilla/mux"

	"cinema.local/db"
	"cinema.local/handlers"
)

func main() {
	// Подключение к базе данных
	db.Connect()

	// Создаём маршруты
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()

	// Роуты фильмов
	api.HandleFunc("/films", handlers.GetFilms).Methods("GET")
	api.HandleFunc("/films", handlers.CreateFilm).Methods("POST")
	api.HandleFunc("/films/{id}", handlers.DeleteFilm).Methods("DELETE")

	// Роуты кинотеатров
	api.HandleFunc("/cinemas", handlers.GetCinemas).Methods("GET")
	api.HandleFunc("/cinemas", handlers.CreateCinema).Methods("POST")

	// Роуты сеансов
	api.HandleFunc("/sessions", handlers.GetSessions).Methods("GET")
	api.HandleFunc("/sessions", handlers.CreateSession).Methods("POST")

	// Роуты аналитики
	api.HandleFunc("/repertoire", handlers.RepertoireByCinema).Methods("GET")
	api.HandleFunc("/cinemas-by-genre", handlers.CinemasByGenre).Methods("GET")
	api.HandleFunc("/session-info", handlers.SessionInfo).Methods("GET")
	api.HandleFunc("/films-by-director", handlers.FilmsByDirector).Methods("GET")

	// Включаем CORS
	corsRouter := gorillahandlers.CORS(
		gorillahandlers.AllowedOrigins([]string{"*"}),
		gorillahandlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"}),
		gorillahandlers.AllowedHeaders([]string{"Content-Type"}),
	)(router)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsRouter))

}
