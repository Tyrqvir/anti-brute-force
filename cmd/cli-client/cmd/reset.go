package cmd

import (
	"context"
	"log"
	"strconv"

	"github.com/Tyrqvir/anti-brute-force/internal/client/grpc"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
	"github.com/spf13/cobra"
)

// ResetCmd represents the add command.
var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset bucket",
	Run: func(cmd *cobra.Command, args []string) {
		client := grpc.NewClient(cmd.PersistentFlags().Lookup("grpcAddress").Value.String())
		ctx := context.Background()

		ui64, err := strconv.ParseUint(cmd.PersistentFlags().Lookup("ip").Value.String(), 10, 32)
		if err != nil {
			log.Fatalf("can't transform string to uint: %v", err)
		}

		_, err = client.AntiBruteForceServiceClient.ResetBucket(ctx, &api.ResetBucketRequest{
			Login: cmd.PersistentFlags().Lookup("login").Value.String(),
			Ip:    uint32(ui64),
		})
		if err != nil {
			log.Fatalf("error on reset bucket: %v", err)
		}
	},
}

func init() {
	ResetCmd.PersistentFlags().StringP("login", "l", "", "login for reset bucket")
	ResetCmd.PersistentFlags().Uint32P("ip", "i", 0, "ip for reset bucket")
	ResetCmd.PersistentFlags().StringP(
		"grpcAddress", "a", "", "grpcAddress",
	)
}
