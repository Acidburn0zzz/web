package http

import (
	"fmt"
	"github.com/bradfitz/http2"
	"github.com/gorilla/mux"
	"log"
	h "net/http"
	"os"
)

// Server encapuslates some common shit
type Server struct {
	Server h.Server
	Router *mux.Router

	port string

	certPath string
	keyPath  string
}

// NewServer creates a new Server
func NewServer(port uint16) *Server {
	s := &Server{}

	certRoot := "/etc/ssl/zqz/zqzca"
	s.certPath = fmt.Sprintf("%s.crt", certRoot)
	s.keyPath = fmt.Sprintf("%s.key", certRoot)
	s.port = fmt.Sprintf(":%d", port)

	s.Router = mux.NewRouter().StrictSlash(true)

	s.Server = h.Server{
		Addr:    s.port,
		Handler: s.Router,
	}

	http2.ConfigureServer(&s.Server, nil)

	return s
}

// Listen starts a HTTP2 server
func (s Server) Listen() {
	if s.certificatesExist() {
		log.Printf("Starting Server. Listening on port :%s\n", s.port)
		s.Server.ListenAndServeTLS(s.certPath, s.keyPath)
	}
}

func (s Server) certificatesExist() bool {
	if _, err := os.Stat(s.certPath); os.IsNotExist(err) {
		log.Printf("Failed to find certificate: (%s)\n", s.certPath)
		return false
	}

	if _, err := os.Stat(s.keyPath); os.IsNotExist(err) {
		log.Printf("Failed to find key: (%s)\n", s.keyPath)
		return false
	}

	return true
}
