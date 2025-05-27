package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"cinema.local/db" // импорт используется
)

func TestCreateFilm(t *testing.T) {
	// Подключение к БД обязательно перед запросами
	db.Connect()

	body := []byte(`{
		"title": "Тестовый фильм",
		"director": "Тестовый режиссёр",
		"genre": "Драма",
		"studio": "Test Studio"
	}`)

	req, err := http.NewRequest("POST", "/api/films", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateFilm)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("expected status 201, got %v", status)
	}
}
