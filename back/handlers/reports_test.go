package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"cinema.local/db"
)

func TestRepertoireByCinema(t *testing.T) {
	db.Connect()

	req, _ := http.NewRequest("GET", "/api/repertoire?cinema=Нева", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(RepertoireByCinema)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %v", rr.Code)
	}
}

func TestCinemasByGenre(t *testing.T) {
	db.Connect()

	req, _ := http.NewRequest("GET", "/api/cinemas-by-genre?genre=комедия", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CinemasByGenre)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %v", rr.Code)
	}
}

func TestSessionInfo(t *testing.T) {
	db.Connect()

	req, _ := http.NewRequest("GET", "/api/session-info?id=1", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(SessionInfo)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %v", rr.Code)
	}
}

func TestFilmsByDirector(t *testing.T) {
	db.Connect()

	req, _ := http.NewRequest("GET", "/api/films-by-director?name=Тестовый режиссёр", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(FilmsByDirector)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %v", rr.Code)
	}
}

func TestFilmsByDirector_MissingParam(t *testing.T) {
	db.Connect()

	req, _ := http.NewRequest("GET", "/api/films-by-director", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(FilmsByDirector)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request, got %v", rr.Code)
	}
}
