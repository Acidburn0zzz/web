package api

import (
	"encoding/json"
	"fmt"
	_ "github.com/davecgh/go-spew/spew"
	. "net/http"
	"strings"
	"testing"
)

func req(method string, url string, jsonRequest string) (*Response, error) {
	params := strings.NewReader(jsonRequest)
	fullURL := fmt.Sprintf("%s%s", server.URL, url)
	req, err := NewRequest("POST", fullURL, params)

	if err != nil {
		return nil, err
	}

	return DefaultClient.Do(req)
}

func expectStatus(t *testing.T, r *Response, expectedCode int) {
	if r.StatusCode != expectedCode {
		t.Error("Should have received status", expectedCode, "got", r.StatusCode)
	}
}

func expectContentType(t *testing.T, r *Response, expectedContentType string) {
	contentType := r.Header.Get("Content-Type")

	if contentType != expectedContentType {
		t.Error("Should have received Content-Type", expectedContentType, "got", contentType)
	}
}

func TestUserController_CreateUserSuccess(t *testing.T) {
	params := `{"first_name":"John", "last_name":"Carmack", "email":"j@id.com"}`
	res, err := req("POST", "/", params)

	if err != nil {
		t.Error(err)
		return
	}

	expectStatus(t, res, 201)
	expectContentType(t, res, "application/json")
}

func TestUserController_CreateUserMalformedJSON(t *testing.T) {
	params := `{"users}`

	res, err := req("POST", "/", params)

	if err != nil {
		t.Error(err)
		return
	}

	expectStatus(t, res, 500)
	expectContentType(t, res, "application/json")
}

func TestUserController_CreateUserInvalid(t *testing.T) {
	params := `{}`

	res, err := req("POST", "/", params)

	if err != nil {
		t.Error(err)
	}

	expectStatus(t, res, 400)
	expectContentType(t, res, "application/json")

	e := &ErrorJSON{}
	if err := json.NewDecoder(res.Body).Decode(e); err != nil {
		t.Error("Failed to decode JSON", err.Error())
	}

	if len(e.Errors) != 3 {
		t.Error(e.Errors)
		t.Error("Should have 3 validation errors. Got", len(e.Errors))
	}
}
