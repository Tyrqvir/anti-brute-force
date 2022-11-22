package cmd //nolint:dupl

import (
	"context"
	"fmt"
	"log"

	"github.com/Tyrqvir/anti-brute-force/internal/client/grpc"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the remove command.
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove from black/white list",
	Run: func(cmd *cobra.Command, args []string) {
		client := grpc.NewClient(globalCfg)
		ctx := context.Background()

		switch cmd.PersistentFlags().Lookup("kindOfList").Value.String() {
		case blackListKey:
			_, err := client.AntiBruteForceServiceClient.RemoveFromBlackList(
				ctx, &api.ListCases{SubNetwork: cmd.PersistentFlags().Lookup("network").Value.String()},
			)
			if err != nil {
				log.Fatalf("error on remove from black list: %v", err)
			}
		case whiteListKey:
			_, err := client.AntiBruteForceServiceClient.RemoveFromWhiteList(
				ctx, &api.ListCases{SubNetwork: cmd.PersistentFlags().Lookup("network").Value.String()},
			)
			if err != nil {
				log.Fatalf("error on remove from white list: %v", err)
			}
		}
	},
}

func init() {
	RemoveCmd.PersistentFlags().StringP(
		"network", "n", "", "network",
	)
	RemoveCmd.PersistentFlags().StringP(
		"kindOfList", "k", blackListKey, fmt.Sprintf("%s or %s", blackListKey, whiteListKey),
	)
}
