package cmd

import (
	"log"
	"os"

	"github.com/Tyrqvir/anti-brute-force/internal/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile      string
	globalCfg    *config.Config
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

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	rootCmd.Flags().StringVarP(&cfgFile, "cfgFile", "c", "../../configs/config.toml", "Path to config file")

	cfg, err := config.NewConfig(cfgFile)
	if err != nil {
		log.Fatalf("can't init config: %v", err)
	}

	globalCfg = cfg
}
