package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"cinema.local/db"
)

func TestLoginUser_Valid(t *testing.T) {
	db.Connect()

	body := []byte(`{
		"username": "testuser123",
		"password": "1234"
	}`)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %v", rr.Code)
	}
}
func TestRegisterUser(t *testing.T) {
	db.Connect()

	body := []byte(`{
		"username": "testuser123",
		"password": "1234"
	}`)

	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Register)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated && rr.Code != http.StatusConflict {
		t.Errorf("expected 201 Created or 409 Conflict, got %v", rr.Code)
	}
}
