package api

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// User API
// POST UserCreate - Create a User

// Response codes
// 201 if succeeded.
// 400 if bad data

// JSON
// {
//   "id": "uuid"
//   "first" : "John",
//   "last" : "Carmack",
//	 "address": "Somewhere in TX",
//   "phone": "+123 123 1234",
//   "email": "johnc@idsoftware.com".
//   "apikey": null,
// }
// {
//   "error": "no password or whatever"
// }
func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	u := User{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Printf("Error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	spew.Dump(u)

	if u.Valid() {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func usersIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to the Dashboard")
}
