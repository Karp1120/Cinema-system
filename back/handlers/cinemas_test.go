package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"cinema.local/db"
)

func TestCreateCinema_Success(t *testing.T) {
	db.Connect()

	body := []byte(`{
		"name": "Киномир",
		"address": "г. Москва",
		"halls": 3
	}`)

	req, _ := http.NewRequest("POST", "/api/cinemas", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateCinema)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %v", rr.Code)
	}
}

func TestCreateCinema_EmptyName(t *testing.T) {
	db.Connect()

	body := []byte(`{
		"name": "",
		"address": "г. Москва",
		"halls": 2
	}`)

	req, _ := http.NewRequest("POST", "/api/cinemas", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateCinema)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request, got %v", rr.Code)
	}
}

func TestGetCinemas(t *testing.T) {
	db.Connect()

	req, _ := http.NewRequest("GET", "/api/cinemas", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetCinemas)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %v", rr.Code)
	}
}
