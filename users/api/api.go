package api

import (
	"github.com/zqzca/web/util/app"
	"upper.io/db"
)

var database db.Database

// UserAPI foo
type UserAPI struct {
	App *app.App
}

// Database is the current DB
func (api *UserAPI) Database() db.Database {
	return *api.App.Database
}

// NewUserAPI derp
func NewUserAPI(dbname string, dbuser string, port int) *UserAPI {
	api := &UserAPI{
		App: app.NewApp(port),
	}

	api.App.AddDatabase(dbname, dbuser)
	api.App.AddMigrations("foo")
	api.App.AddRoute("GET", "/{id}", usersRead)
	api.App.AddRoute("GET", "/", usersIndex)
	api.App.AddRoute("POST", "/", UserCreate)

	database = *api.App.Database

	return api
}
