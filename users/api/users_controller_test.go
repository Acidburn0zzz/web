package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestUserController_CreateUserSuccess(t *testing.T) {
	params := buildValidUserJSON()
	res, err := req("POST", "/", params)

	if err != nil {
		t.Error(err)
		return
	}

	expectStatus(t, res, 201)
	expectContentType(t, res, "application/json")

	u := decodeUser(t, res.Body)

	expectNotNil(t, u.ID)
	expectPresent(t, "ID", u.ID)
	expectEqual(t, u.FirstName, "John")
	expectEqual(t, u.LastName, "Carmack")
	expectEqual(t, u.Email, "johnc@idsoftware.com")
}

func TestUserController_CreateUserPersisted(t *testing.T) {
	truncateUsers()
	params := buildValidUserJSON()
	res, err := req("POST", "/", params)

	if err != nil {
		t.Error(err)
		return
	}

	expectStatus(t, res, 201)
	u := decodeUser(t, res.Body)

	UserCount()

	path := fmt.Sprintf("/%s", u.ID)

	spew.Dump(path)

	res, err = req("GET", path, "")

	if err != nil {
		t.Error(err)
		return
	}

	expectStatus(t, res, 200)
	u2 := decodeUser(t, res.Body)

	expectEqual(t, u.ID, u2.ID)
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

	expectStatus(t, res, http.StatusNotAcceptable)
	expectContentType(t, res, "application/json")

	e := &ErrorJSON{}
	if err := json.NewDecoder(res.Body).Decode(e); err != nil {
		t.Error("Failed to decode JSON", err.Error())
	}

	if len(e.Errors) != 4 {
		t.Error(e.Errors)
		t.Error("Should have 4 validation errors. Got", len(e.Errors))
	}
}
