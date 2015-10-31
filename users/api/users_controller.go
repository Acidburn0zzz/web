package api

import (
	"encoding/json"
	// fuck off
	_ "github.com/davecgh/go-spew/spew"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// ErrorJSON is a simple blew
type ErrorJSON struct {
	Errors []string `json:"errors"`
}

// UserCreate POST - Create a User
//
// Response codes
// 201 if succeeded
// 400 if bad data
//
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
//   "error": [
//	   "no password or whatever",
//     "invalid email address",
//	 ]
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

	userResponse(w, u)
}

func userResponse(w http.ResponseWriter, u User) {
	if u.Valid() {
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(u); err != nil {
			log.Printf("Error: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)

		e := &ErrorJSON{Errors: u.errors}
		if err := json.NewEncoder(w).Encode(e); err != nil {
			log.Printf("Error: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func usersIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to the Dashboard")
}
