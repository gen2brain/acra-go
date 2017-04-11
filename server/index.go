package server

import (
	"fmt"
	"net/http"
)

// Index struct
type Index struct {
}

// NewIndex returns new Index
func NewIndex() *Index {
	return &Index{}
}

// ServeHTTP handles requests on incoming connections
func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "HEAD" {
		msg := fmt.Sprintf("405 Method Not Allowed (%s)", r.Method)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	data, err := Asset("app/index.html")
	if err != nil {
		msg := fmt.Sprintf("500 Internal Server Error (%s)", err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
