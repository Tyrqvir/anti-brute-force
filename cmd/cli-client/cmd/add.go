package cmd //nolint:dupl

import (
	"context"
	"fmt"
	"log"

	"github.com/Tyrqvir/anti-brute-force/internal/client/grpc"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
	"github.com/spf13/cobra"
)

// AddCmd represents the add command.
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add network to black/white list",
	Run: func(cmd *cobra.Command, args []string) {
		client := grpc.NewClient(globalCfg)
		ctx := context.Background()

		switch cmd.PersistentFlags().Lookup("kindOfList").Value.String() {
		case blackListKey:
			_, err := client.AntiBruteForceServiceClient.AddToBlackList(
				ctx, &api.ListCases{SubNetwork: cmd.PersistentFlags().Lookup("network").Value.String()},
			)
			if err != nil {
				log.Fatalf("error on add to black list: %v", err)
			}
		case whiteListKey:
			_, err := client.AntiBruteForceServiceClient.AddToWhiteList(
				ctx, &api.ListCases{SubNetwork: cmd.PersistentFlags().Lookup("network").Value.String()},
			)
			if err != nil {
				log.Fatalf("error on add to white list: %v", err)
			}
		}
	},
}

func init() {
	AddCmd.PersistentFlags().StringP(
		"network", "n", "", "network",
	)
	AddCmd.PersistentFlags().StringP(
		"kindOfList", "k", blackListKey, fmt.Sprintf("%s or %s", blackListKey, whiteListKey),
	)
}
