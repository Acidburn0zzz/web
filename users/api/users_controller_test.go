package api

import (
	. "net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserController_CreateUserSuccess(t *testing.T) {
	params := strings.NewReader(
		`{"first_name":"John", "last_name":"Carmack", "email":"johnc@idsoftware.com"}`,
	)

	req, err := NewRequest("POST", "http://zqz.ca/", params)

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	UserCreate(w, req)

	if w.Code != 201 {
		t.Error("Should have received status 201, got", w.Code)
	}

	contentType := w.HeaderMap.Get("Content-Type")

	if contentType != "application/json" {
		t.Error("Should have received content type application/json got", contentType)
	}
}

func TestUserController_CreateUserFailure(t *testing.T) {
	params := strings.NewReader(
		`{}`,
	)

	req, err := NewRequest("POST", "http://zqz.ca/", params)

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	UserCreate(w, req)

	if w.Code != 400 {
		t.Error("Should have received status 400, got", w.Code)
	}

	contentType := w.HeaderMap.Get("Content-Type")

	if contentType != "application/json" {
		t.Error("Should have received content type application/json got", contentType)
	}
}
