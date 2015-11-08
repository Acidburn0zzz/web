package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func buildValidUser() *User {
	return &User{
		FirstName: "John",
		LastName:  "Carmack",
		Address:   "somewhere in texas",
		Username:  "jc",
		Phone:     "+123 123 1234",
		Email:     "johnc@idsoftware.com",
	}
}

func buildValidUserJSON() string {
	return buildValidUser().String()
}

func req(method string, url string, jsonRequest string) (*http.Response, error) {
	params := strings.NewReader(jsonRequest)
	fullURL := fmt.Sprintf("%s%s", server.URL, url)
	req, err := http.NewRequest(method, fullURL, params)

	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

func expectStatus(t *testing.T, r *http.Response, expectedCode int) {
	if r.StatusCode != expectedCode {
		t.Error("Should have received status", expectedCode, "got", r.StatusCode)
	}
}

func expectContentType(t *testing.T, r *http.Response, expectedContentType string) {
	contentType := r.Header.Get("Content-Type")

	if contentType != expectedContentType {
		t.Error("Should have received Content-Type", expectedContentType, "got", contentType)
	}
}

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

func expectNil(t *testing.T, o interface{}) {
	if !isNil(o) {
		t.Error("Expected nil but got non nil value")
	}
}

func expectEqual(t *testing.T, s interface{}, e interface{}) {
	if s != e {
		t.Error("Expected", s, "got", e)
	}
}

func expectNotNil(t *testing.T, o interface{}) {
	if isNil(o) {
		t.Error("Expected", o, "not to be nil")
	}
}

func expectPresent(t *testing.T, name string, o interface{}) {
	if o == nil {
		t.Error("Expected", o, "to be present but was nil")
	}

	switch v := o.(type) {
	case string:
		if len(v) == 0 {
			t.Error("Expected", name, "to be present but was empty string")
		}
	default:
		t.Error("Expected", name, "to be present but was unknown type")
	}
}

func decodeUser(t *testing.T, body io.Reader) *User {
	u := &User{}

	if err := json.NewDecoder(body).Decode(u); err != nil {
		t.Error("Failed to decode JSON", err.Error())
	}

	return u
}
