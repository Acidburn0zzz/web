package app

import (
	"fmt"
	"github.com/bradfitz/http2"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"upper.io/db"
	"upper.io/db/postgresql"
)

// App encapuslates some common shit
type App struct {
	Server   *http.Server
	Router   *mux.Router
	Database *db.Database

	port     string
	certPath string
	keyPath  string
}

// NewApp creates a new App
func NewApp(port int) *App {
	a := &App{}

	certRoot := "/etc/ssl/zqz/zqzca"
	a.certPath = fmt.Sprintf("%s.crt", certRoot)
	a.keyPath = fmt.Sprintf("%s.key", certRoot)
	a.port = fmt.Sprintf(":%d", port)

	a.Router = mux.NewRouter().StrictSlash(true)

	a.Server = &http.Server{
		Addr:    a.port,
		Handler: a.Router,
	}

	http2.ConfigureServer(a.Server, nil)

	return a
}

// AddRoute adds a route
func (a App) AddRoute(method string, path string, route http.HandlerFunc) {
	a.Router.Methods(method).Path(path).Handler(route)
}

// Listen starts a HTTP2 server
func (a App) Listen() {
	if a.certificatesExist() {
		log.Printf("Starting Application. Listening on port :%s\n", a.port)
		a.Server.ListenAndServeTLS(a.certPath, a.keyPath)
	}
}

// AddDatabase connects to a psql db with the given name.
func (a App) AddDatabase(name string, user string) {
	settings := postgresql.ConnectionURL{
		Database: name,
		User:     user,
	}

	database, err := db.Open(postgresql.Adapter, settings)

	if err != nil {
		log.Fatalf("Failed to connect to database: %s with user %s - %s\n", name, user, err.Error())
	}

	if err = database.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %s - %s\n", name, err.Error())
	}

	a.Database = &database
}

func (a App) certificatesExist() bool {
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
