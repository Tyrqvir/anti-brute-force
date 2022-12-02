//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Tyrqvir/anti-brute-force/internal/config"
	"github.com/Tyrqvir/anti-brute-force/internal/limiter/lickybucket"
	"github.com/Tyrqvir/anti-brute-force/internal/logger"
	"github.com/Tyrqvir/anti-brute-force/internal/server/grpc"
	"github.com/Tyrqvir/anti-brute-force/internal/storage"
	"github.com/Tyrqvir/anti-brute-force/internal/storage/redis"
	"github.com/Tyrqvir/anti-brute-force/internal/usercases"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
	"github.com/google/wire"
)

func setupWire(config *config.Config, logger logger.Logger) (*grpc.Server, error) {
	wire.Build(
		wire.Bind(new(storage.Storage), new(*redis.StorageRedis)),
		wire.Bind(new(api.AntiBruteForceServiceServer), new(*usercases.AntiBruteForceCase)),
		usercases.NewUseCase,
		lickybucket.New,
		redis.New,
		grpc.New,
	)
	return &grpc.Server{}, nil
}
