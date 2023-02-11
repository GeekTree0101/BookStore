package cmd

import (
	"github.com/Geektree0101/bookstore/app"
	"github.com/spf13/cobra"
)

var (
	debugCmd = &cobra.Command{
		Use:   "debug",
		Short: "run debug server",
		Run: func(cmd *cobra.Command, args []string) {
			app.RunServer()
		},
	}
)

func init() {
	rootCmd.AddCommand(debugCmd)
}
