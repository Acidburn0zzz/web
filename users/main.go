package main

import (
	"github.com/zqzca/web/users/api"
	"github.com/zqzca/web/util/app"
	"net/http"

	"io"
)

func usersIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to the Dashboard")
}

func main() {
	database := "zqz-users-dev"
	dbuser := "dylan"
	port := 9000

	s := app.NewApp(port)
	s.AddDatabase(database, dbuser)
	s.AddRoute("GET", "/", usersIndex)
	s.AddRoute("POST", "/", api.UserCreate)
	s.Listen()
}
