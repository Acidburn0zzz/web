package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"
)

var (
	api    *UserAPI
	server *httptest.Server
	dbconn *sql.DB
)

func init() {
	api = NewUserAPI("zqz-users-test", "dylan", 5000)
	dbconn = api.Database().Driver().(*sql.DB)
	truncateUsers()
	server = httptest.NewServer(api.App.Router)
}

func truncateUsers() {
	dbconn.Query("truncate users;")
}

func transaction(t *testing.T, f func(t *testing.T)) {
	drv := api.Database().Driver().(*sql.DB)
	tx, err := drv.Begin()

	fmt.Println("WTF")

	if err != nil {
		log.Fatalln("failed to make transaction")
	}

	f(t)
	tx.Rollback()
	fmt.Println("AAWTF")
}
