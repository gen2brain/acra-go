package cmd

import (
	"log"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

var restartFlags struct {
	Name string
}

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart acra-go",
	Long:  `The restart ops command restarts the acra-go service`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := service.New(nil, &service.Config{Name: restartFlags.Name})
		if nil != err {
			log.Fatalln(err)
		}
		if err = s.Restart(); nil != err {
			log.Fatalln(err)
		}
	},
}

func init() {
	opsCmd.AddCommand(restartCmd)

	restartCmd.Flags().StringVarP(&restartFlags.Name, `name`, `n`, `acra-go`, `Name of the service`)
}
