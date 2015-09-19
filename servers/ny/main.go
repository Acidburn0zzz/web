package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	key := "v1"
	e := `"` + key + `"`
	w.Header().Set("Etag", e)
	w.Header().Set("Cache-Control", "max-age=2592000") // 30 days

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
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
