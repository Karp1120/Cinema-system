package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"cinema.local/db"
)

func TestCreateSession_Valid(t *testing.T) {
	db.Connect()

	body := []byte(`{
		"cinema_id": 1,
		"film_id": 1,
		"date": "2025-12-31",
		"time": "22:00:00",
		"price": 500,
		"available_seats": 100,
		"hall_number": 99
	}`)

	req, _ := http.NewRequest("POST", "/api/sessions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateSession)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %v", rr.Code)
	}
}
