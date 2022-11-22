package grpc

import (
	"log"

	"github.com/Tyrqvir/anti-brute-force/internal/config"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	AntiBruteForceServiceClient api.AntiBruteForceServiceClient
}

func NewClient(
	config *config.Config,
) *Client {
	conn, err := grpc.Dial(config.GRPC.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error start grpc client on " + config.GRPC.Address)
	}

	return &Client{
		AntiBruteForceServiceClient: api.NewAntiBruteForceServiceClient(conn),
	}
}
