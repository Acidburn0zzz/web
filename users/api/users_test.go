package api

import "testing"

func TestUser_Valid(t *testing.T) {
	t.Parallel()

	// empty user is not valid
	u := &User{}

	if u.Valid() {
		t.Error("Empty user should not be valid.")
	}

	// factory is valid
	u = buildValidUser()

	if !u.Valid() {
		t.Error("buildValidUser should build a valid user.")
	}

	// missing email is not valid
	u.Email = ""

	if u.Valid() {
		t.Error("missing email should be invalid")
	}

	// missing name is not valid
	u = buildValidUser()
	u.FirstName = ""

	if u.Valid() {
		t.Error("missing first name should be invalid")
	}

	u = buildValidUser()
	u.LastName = ""

	if u.Valid() {
		t.Error("missing last name should be invalid")
	}
}

func buildValidUser() *User {
	return &User{
		FirstName: "John",
		LastName:  "Carmack",
		Address:   "somewhere in texas",
		Phone:     "+123 123 1234",
		Email:     "johnc@idsoftware.com",
	}
}
