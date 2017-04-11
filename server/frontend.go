package server

import (
	"fmt"
	"net/http"

	"github.com/gen2brain/acra-go/acra"
	"github.com/gen2brain/acra-go/database"
)

// Frontend struct
type Frontend struct {
	DB database.DB
}

// NewFrontend returns new Frontend
func NewFrontend(db database.DB) *Frontend {
	return &Frontend{db}
}

// ServeHTTP handles requests on incoming connections
func (f *Frontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET", "HEAD":
		key := r.FormValue("id")

		if key != "" {
			report, err := f.DB.Get(key)
			if err != nil {
				msg := fmt.Sprintf("400 Bad Request (%s)", err.Error())
				http.Error(w, msg, http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			encoder := acra.NewEncoder(w)

			err = encoder.Encode(report)
			if err != nil {
				msg := fmt.Sprintf("500 Internal Server Error (%s)", err.Error())
				http.Error(w, msg, http.StatusInternalServerError)
				return
			}
		} else {
			reports, err := f.DB.GetAll()
			if err != nil {
				msg := fmt.Sprintf("400 Bad Request (%s)", err.Error())
				http.Error(w, msg, http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			encoder := acra.NewEncoder(w)

			err = encoder.Encode(reports)
			if err != nil {
				msg := fmt.Sprintf("500 Internal Server Error (%s)", err.Error())
				http.Error(w, msg, http.StatusInternalServerError)
				return
			}
		}
	case "DELETE":
		key := r.FormValue("id")
		if key != "" {
			err := f.DB.Delete(key)
			if err != nil {
				msg := fmt.Sprintf("400 Bad Request (%s)", err.Error())
				http.Error(w, msg, http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
			return
		}

		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	default:
		msg := fmt.Sprintf("405 Method Not Allowed (%s)", r.Method)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}
}
