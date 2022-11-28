package grpc

import (
	"log"

	"github.com/Tyrqvir/anti-brute-force/proto/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	AntiBruteForceServiceClient api.AntiBruteForceServiceClient
}

func NewClient(
	address string,
) *Client {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error start grpc client on " + address)
	}
	return &Client{
		AntiBruteForceServiceClient: api.NewAntiBruteForceServiceClient(conn),
	}
}
