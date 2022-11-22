package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/Tyrqvir/anti-brute-force/internal/config"
	"github.com/Tyrqvir/anti-brute-force/internal/logger"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

type Server struct {
	server  *grpc.Server
	logger  logger.Logger
	address string
}

func New(
	antiBruteForceServiceServer api.AntiBruteForceServiceServer,
	logger logger.Logger,
	config *config.Config,
) *Server {
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(logger.GetInstance()),
		),
	)

	api.RegisterAntiBruteForceServiceServer(server, antiBruteForceServiceServer)

	return &Server{
		logger:  logger,
		server:  server,
		address: config.GRPC.Address,
	}
}

func (s *Server) Start(ctx context.Context) error {
	listenConfig := net.ListenConfig{}

	listener, err := listenConfig.Listen(ctx, "tcp", s.address)
	if err != nil {
		return fmt.Errorf("start grpc server failed: %w", err)
	}

	s.logger.Info("start grpc server on " + s.address)

	return s.server.Serve(listener)
}

func (s *Server) Stop() {
	s.logger.Info("stop grpc server...")

	s.server.GracefulStop()
}
