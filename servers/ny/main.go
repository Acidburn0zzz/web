package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	_ "github.com/zqzca/web/daemon"
	"github.com/zqzca/web/util/app"
)

func handler(w http.ResponseWriter, r *http.Request) {
	key := "v1"
	e := `"` + key + `"`
	// w.Header().Set("Etag", e)
	// w.Header().Set("Cache-Control", "max-age=2592000") // 30 days

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, e) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	data, err := ioutil.ReadFile("index.html")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 Not Found")
		return
	}

	io.WriteString(w, string(data))
}

func main() {
	app := app.NewApp(8080)
	app.AddRoute("GET", "/", handler)
	app.Listen()
}
