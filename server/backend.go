package server

import (
	"fmt"
	"net/http"

	"github.com/gen2brain/acra-go/acra"
	"github.com/gen2brain/acra-go/database"
)

// Backend struct
type Backend struct {
	DB database.DB
}

// NewBackend returns new Backend
func NewBackend(db database.DB) *Backend {
	return &Backend{db}
}

// ServeHTTP handles requests on incoming connections
func (b *Backend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method != "POST" && r.Method != "PUT" {
		msg := fmt.Sprintf("405 Method Not Allowed (%s)", r.Method)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	contentType := r.Header.Get("Content-type")
	if contentType != "application/json" && contentType != "application/x-www-form-urlencoded" {
		msg := fmt.Sprintf("415 Unsupported Media Type (%s)", contentType)
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	report := acra.Report{}
	decoder := acra.NewDecoder(r.Body)

	err = decoder.Decode(&report)
	if err != nil {
		msg := fmt.Sprintf("400 Bad Request (%s)", err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	err = b.DB.Put(report.ReportID, report)
	if err != nil {
		msg := fmt.Sprintf("500 Internal Server Error (%s)", err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
