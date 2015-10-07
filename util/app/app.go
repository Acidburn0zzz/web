package app

import (
	"fmt"
	"github.com/bradfitz/http2"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

// App encapuslates some common shit
type App struct {
	Server *http.Server
	Router *mux.Router

	port string

	certPath string
	keyPath  string
}

// NewApp creates a new App
func NewApp(port uint16) *App {
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
