package api

import (
	"net/http/httptest"
)

var (
	server *httptest.Server
)

func init() {
	api := NewUserAPI("zqz-users-test", "dylan", 5000)
	server = httptest.NewServer(api.App.Router)
}
