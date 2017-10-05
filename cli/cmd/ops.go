package cmd

import (
	"github.com/spf13/cobra"
)

// opsCmd represents the ops command
var opsCmd = &cobra.Command{
	Use:   "ops",
	Short: "Dev-ops related commands",
	Long:  `Ops performs devops functions for acra-go: install, restart, uninstall, etc.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// TODO: Work your own magic here
	// 	fmt.Println("ops called")
	// },
}

func init() {
	RootCmd.AddCommand(opsCmd)
}
