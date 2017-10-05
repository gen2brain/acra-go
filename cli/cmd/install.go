package cmd

import (
	"log"
	"os"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

var installFlags struct {
	Dir  string
	Name string
	User string
}

// installCmd handles system installation of acra-go
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install acra-go as a system service",
	Long:  `Installs acra-go as a system service.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if `` == installFlags.Dir {
			installFlags.Dir, err = os.Getwd()
			if nil != err {
				log.Fatalln(err)
			}
		}
		serviceArgs := make([]string, len(args)+1)
		serviceArgs[0] = `serve`
		if 0 < len(args) {
			copy(serviceArgs[1:], args)
		}
		cfg := service.Config{
			Name:             installFlags.Name,
			DisplayName:      `acra-go service`,
			Description:      `acra-go service for logging and reporting on Android ACRA enabled crash reports`,
			UserName:         installFlags.User,
			Arguments:        serviceArgs,
			Executable:       ``,         // Execute ourselves
			Dependencies:     []string{}, // Not supported on Linux
			WorkingDirectory: installFlags.Dir,
			ChRoot:           ``, // We'l not chroot for now
		}
		s, err := service.New(nil, &cfg)
		if nil != err {
			log.Fatalln(err)
		}
		s.Uninstall() // We uninstall, but ignore any errors that might occur
		if err = s.Install(); nil != err {
			log.Fatalln(err)
		}
	},
}

func init() {
	opsCmd.AddCommand(installCmd)

	opsCmd.Flags().StringVarP(&installFlags.Dir, `dir`, `d`, ``, `Directory of acra-go: default to current working dir`)
	opsCmd.Flags().StringVarP(&installFlags.Name, `name`, `n`, `acra-go`, `Name of the service`)
	opsCmd.Flags().StringVarP(&installFlags.User, `user`, `u`, `acra-go`, `User to run as`)
}
