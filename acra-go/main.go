// +build leveldb

package main

import (
	"flag"
	"log"
	"path/filepath"

	bolt "github.com/gen2brain/acra-go/database/drivers/bolt"
	level "github.com/gen2brain/acra-go/database/drivers/leveldb"
	scribble "github.com/gen2brain/acra-go/database/drivers/scribble"

	"github.com/gen2brain/acra-go/database"
	"github.com/gen2brain/acra-go/server"
)

func main() {
	srv := server.NewServer(nil)

	flag.StringVar(&srv.Bind, "bind-addr", ":55000", "Bind address")
	flag.StringVar(&srv.DatabaseDir, "database-dir", ".", "Path to database directory")
	flag.StringVar(&srv.HtpasswdBackend, "htpasswd-backend", "", "Path to htpasswd file, if empty backend auth is disabled")
	flag.StringVar(&srv.HtpasswdFrontend, "htpasswd-frontend", "", "Path to htpasswd file, if empty frontend auth is disabled")
	flag.IntVar(&srv.ReadTimeout, "read-timeout", 5, "Read timeout (seconds)")
	flag.IntVar(&srv.WriteTimeout, "write-timeout", 15, "Write timeout (seconds)")

	engine := flag.String(`engine`, `level`, `Database engine to use (bolt|level|scribble)`)

	flag.Parse()

	suffix := ``

	var db database.DB
	switch *engine {
	case `bolt`:
		db = bolt.NewDB()
		suffix = `.db`
	case `level`:
		db = level.NewDB()
	case `scribble`:
		db = scribble.NewDB()
	default:
		log.Fatalf(`Unrecognized engine '%s': please use one of bolt|level|scribble`, *engine)
	}
	srv.Database = db
	db.Open(filepath.Join(srv.DatabaseDir, "reports"+suffix))
	defer db.Close()

	srv.ListenAndServe()
}
