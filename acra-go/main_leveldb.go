// +build leveldb

package main

import (
	"flag"
	"path/filepath"

	"github.com/gen2brain/acra-go/database/drivers/leveldb"
	"github.com/gen2brain/acra-go/server"
)

func main() {
	db := driver.NewDB()
	srv := server.NewServer(db)

	flag.StringVar(&srv.Bind, "bind-addr", ":55000", "Bind address")
	flag.StringVar(&srv.DatabaseDir, "database-dir", ".", "Path to database directory")
	flag.StringVar(&srv.HtpasswdBackend, "htpasswd-backend", "", "Path to htpasswd file, if empty backend auth is disabled")
	flag.StringVar(&srv.HtpasswdFrontend, "htpasswd-frontend", "", "Path to htpasswd file, if empty frontend auth is disabled")
	flag.IntVar(&srv.ReadTimeout, "read-timeout", 5, "Read timeout (seconds)")
	flag.IntVar(&srv.WriteTimeout, "write-timeout", 15, "Write timeout (seconds)")
	flag.Parse()

	db.Open(filepath.Join(srv.DatabaseDir, "reports"))
	defer db.Close()

	srv.ListenAndServe()
}
