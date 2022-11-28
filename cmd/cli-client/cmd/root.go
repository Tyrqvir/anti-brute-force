package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	whiteListKey = "white"
	blackListKey = "black"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "cli-client",
	Short: "The grpc client for anti brute force service",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(RemoveCmd)
	rootCmd.AddCommand(ResetCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
