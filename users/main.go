package main

import (
	"github.com/zqzca/web/users/api"
	"github.com/zqzca/web/util/app"
	"net/http"

	"io"
)

func main() {
	api := NewUserAPI("zqz-users-dev", "dylan", 5000)
	api.App.Listen()
}
