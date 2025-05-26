package main

import (
	"log"
	"net/http"

	"cinema.local/db"
	"cinema.local/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db.Connect()

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

	// Аналитика
	api.HandleFunc("/repertoire", handlers.RepertoireByCinema).Methods("GET")
	api.HandleFunc("/cinemas-by-genre", handlers.CinemasByGenre).Methods("GET")
	api.HandleFunc("/session-info", handlers.SessionInfo).Methods("GET")
	api.HandleFunc("/films-by-director", handlers.FilmsByDirector).Methods("GET")

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
