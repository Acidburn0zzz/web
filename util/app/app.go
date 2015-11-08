package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bradfitz/http2"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"upper.io/db"
	"upper.io/db/postgresql"
)

var database db.Database

// App encapuslates some common shit
type App struct {
	Server   *http.Server
	Router   *mux.Router
	Database *db.Database

	port     string
	certPath string
	keyPath  string

	hasMigrations bool
}

// NewApp creates a new App
func NewApp(port int) *App {
	a := &App{}

	certRoot := "/etc/ssl/zqz/zqzca"
	a.certPath = fmt.Sprintf("%s.crt", certRoot)
	a.keyPath = fmt.Sprintf("%s.key", certRoot)
	a.port = fmt.Sprintf(":%d", port)

	a.Router = mux.NewRouter().StrictSlash(true)

	chain := alice.New().Then(a.Router)

	a.Server = &http.Server{
		Addr:    a.port,
		Handler: chain,
	}

	http2.ConfigureServer(a.Server, nil)

	return a
}

// AddRoute adds a route
func (a *App) AddRoute(method string, path string, route http.HandlerFunc) {
	// a.Router.Methods(method).Path(path).Handler(
	// 	func(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	// 		route(w, r)
	// 	},
	// )
	a.Router.Methods(method).Path(path).Handler(route)
}

// Listen starts a HTTP2 server
func (a *App) Listen() {
	if a.certificatesExist() {
		log.Printf("Starting Application. Listening on port :%s\n", a.port)
		a.Server.ListenAndServeTLS(a.certPath, a.keyPath)
	}
}

// AddDatabase connects to a psql db with the given name.
func (a *App) AddDatabase(name string, user string) {
	settings := postgresql.ConnectionURL{
		Database: name,
		User:     user,
	}

	var err error
	database, err = db.Open(postgresql.Adapter, settings)

	if err != nil {
		log.Fatalf("Failed to connect to database: %s with user %s - %s\n", name, user, err.Error())
	}

	if err = database.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %s - %s\n", name, err.Error())
	}

	a.Database = &database
}

func (a *App) certificatesExist() bool {
	if _, err := os.Stat(a.certPath); os.IsNotExist(err) {
		log.Printf("Failed to find certificate: (%s)\n", a.certPath)
		return false
	}

	if _, err := os.Stat(a.keyPath); os.IsNotExist(err) {
		log.Printf("Failed to find key: (%s)\n", a.keyPath)
		return false
	}

	return true
}
