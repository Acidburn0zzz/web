package api

import (
	"github.com/zqzca/web/util/app"
)

// UserAPI foo
type UserAPI struct {
	App *app.App
}

// NewUserAPI derp
func NewUserAPI(dbname string, dbuser string, port int) *UserAPI {
	api := &UserAPI{
		App: app.NewApp(port),
	}

	api.App.AddDatabase(dbname, dbuser)
	api.App.AddRoute("GET", "/", usersIndex)
	api.App.AddRoute("POST", "/", UserCreate)

	return api
}
