package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "bookStore",
		Short: "BookStore example",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
