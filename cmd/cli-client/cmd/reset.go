package cmd

import (
	"context"
	"log"

	"github.com/Tyrqvir/anti-brute-force/internal/client/grpc"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
	"github.com/spf13/cobra"
)

// ResetCmd represents the add command.
var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset bucket",
	Run: func(cmd *cobra.Command, args []string) {
		client := grpc.NewClient(globalCfg)
		ctx := context.Background()

		_, err := client.AntiBruteForceServiceClient.ResetBucket(ctx, &api.ResetBucketRequest{
			Login: cmd.PersistentFlags().Lookup("login").Value.String(),
			Ip:    cmd.PersistentFlags().Lookup("ip").Value.String(),
		})
		if err != nil {
			log.Fatalf("error on reset bucket: %v", err)
		}
	},
}

func init() {
	ResetCmd.PersistentFlags().StringP("login", "l", "", "login for reset bucket")
	ResetCmd.PersistentFlags().StringP("ip", "i", "", "password for reset bucket")
}
