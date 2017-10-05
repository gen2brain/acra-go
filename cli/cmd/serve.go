package cmd

import (
	"log"
	"path/filepath"

	bolt "github.com/gen2brain/acra-go/database/drivers/bolt"
	level "github.com/gen2brain/acra-go/database/drivers/leveldb"
	scribble "github.com/gen2brain/acra-go/database/drivers/scribble"

	"github.com/gen2brain/acra-go/database"
	"github.com/gen2brain/acra-go/server"

	"github.com/spf13/cobra"
)

var serveFlags struct {
	BindAddr         string
	DatabaseDir      string
	HtpasswdBackend  string
	HtpasswdFrontend string
	ReadTimeout      int
	WriteTimeout     int
	Engine           string
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run acra-go webserver",
	Long: `The acra-go webserver captures reports from ACRA enabled Android clients,
	provides a UI for viewing the reports received.`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.NewServer(nil)
		srv.Bind = serveFlags.BindAddr
		srv.DatabaseDir = serveFlags.DatabaseDir
		srv.HtpasswdBackend = serveFlags.HtpasswdBackend
		srv.HtpasswdFrontend = serveFlags.HtpasswdFrontend
		srv.ReadTimeout = serveFlags.ReadTimeout
		srv.WriteTimeout = serveFlags.WriteTimeout

		suffix := ``

		var db database.DB
		switch serveFlags.Engine {
		case `bolt`:
			db = bolt.NewDB()
			suffix = `.db`
		case `level`:
			db = level.NewDB()
		case `scribble`:
			db = scribble.NewDB()
		default:
			log.Fatalf(`Unrecognized engine '%s': please use one of bolt|level|scribble`, serveFlags.Engine)
		}
		srv.Database = db
		db.Open(filepath.Join(srv.DatabaseDir, "reports"+suffix))
		defer db.Close()

		srv.ListenAndServe()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&serveFlags.BindAddr, `bind-addr`, `p`, `:55000`, `Bind address`)
	serveCmd.Flags().StringVarP(&serveFlags.DatabaseDir, `database-dir`, `d`, `.`, `Path to database directory`)
	serveCmd.Flags().StringVarP(&serveFlags.HtpasswdBackend, `htpasswd-backend`, `b`, ``, `Path to htpasswd file, if empty backend auth is disabled`)
	serveCmd.Flags().StringVarP(&serveFlags.HtpasswdFrontend, `htpasswd-frontend`, `f`, ``, `Path to htpasswd file, if empty frontend auth is disabled`)
	serveCmd.Flags().IntVarP(&serveFlags.ReadTimeout, `read-timeout`, `r`, 5, `Read timeout (seconds)`)
	serveCmd.Flags().IntVarP(&serveFlags.WriteTimeout, `write-timeout`, `w`, 5, `Write timeout (seconds)`)
	serveCmd.Flags().StringVarP(&serveFlags.Engine, `engine`, `e`, `level`, `Database engine to use (bolt|level|scribble)`)
}
