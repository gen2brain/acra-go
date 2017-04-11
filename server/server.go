package server

//go:generate go-bindata -nocompress -pkg server assets/... app/...

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/abbot/go-http-auth"
	"github.com/elazarl/go-bindata-assetfs"

	"github.com/gen2brain/acra-go/database"
)

const (
	// Name is server name
	Name = "acra-go"
	// Version is server version
	Version = "1.0"
)

// Server struct
type Server struct {
	Bind             string
	Database         database.DB
	DatabaseDir      string
	HtpasswdBackend  string
	HtpasswdFrontend string
	ReadTimeout      int
	WriteTimeout     int
}

// NewServer returns new Server
func NewServer(db database.DB) *Server {
	s := &Server{}
	s.Database = db
	return s
}

// ListenAndServe listens on the TCP address and serves requests
func (s *Server) ListenAndServe() {
	var authBackend, authFrontend *auth.BasicAuth
	realm := fmt.Sprintf("%s/%s", Name, Version)

	if s.HtpasswdBackend != "" {
		if _, err := os.Stat(s.HtpasswdBackend); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		authBackend = auth.NewBasicAuthenticator(realm, auth.HtpasswdFileProvider(s.HtpasswdBackend))
	}

	if s.HtpasswdFrontend != "" {
		if _, err := os.Stat(s.HtpasswdFrontend); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		authFrontend = auth.NewBasicAuthenticator(realm, auth.HtpasswdFileProvider(s.HtpasswdFrontend))
	}

	http.Handle("/view", newAuthHandler(NewFrontend(s.Database), authFrontend))
	http.Handle("/send", newAuthHandler(NewBackend(s.Database), authBackend))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		data, err := Asset("assets/favicon.ico")
		if err == nil {
			w.Write(data)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User-agent: *\nDisallow: /"))
	})

	http.Handle("/assets/", http.FileServer(&assetfs.AssetFS{Asset, AssetDir, AssetInfo, ""}))

	http.Handle("/", newAuthHandler(NewIndex(), authFrontend))

	srv := &http.Server{
		ReadTimeout:  time.Duration(s.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.WriteTimeout) * time.Second,
	}

	listener, err := net.Listen("tcp4", s.Bind)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	srv.Serve(listener)
}

// newAuthHandler wraps handler and checks auth
func newAuthHandler(handler http.Handler, authenticator *auth.BasicAuth) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				fmt.Fprintf(os.Stderr, "Recovered: %v\n", rec)
			}
		}()

		w.Header().Set("Server", fmt.Sprintf("%s/%s", Name, Version))

		if authenticator != nil {
			w.Header().Set("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", authenticator.Realm))
			if authenticator.CheckAuth(r) == "" {
				http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		handler.ServeHTTP(w, r)
	})
}
