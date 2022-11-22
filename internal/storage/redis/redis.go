package redis

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/Tyrqvir/anti-brute-force/internal/config"
	"github.com/Tyrqvir/anti-brute-force/internal/logger"
	"github.com/Tyrqvir/anti-brute-force/internal/storage"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

const (
	BlackListKey = "blacklist"
	WhiteListKey = "whitelist"
)

type StorageRedis struct {
	storage.Storage

	rdb     *redis.Client
	logger  logger.Logger
	config  *config.Config
	limiter *redis_rate.Limiter
}

func New(config *config.Config, logger logger.Logger) *StorageRedis {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.Redis.DSN,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logger.Errorf("error connect to redis: %v", err)
	}

	limiter := redis_rate.NewLimiter(rdb)

	return &StorageRedis{
		rdb:     rdb,
		logger:  logger,
		config:  config,
		limiter: limiter,
	}
}

func (s *StorageRedis) FlushAll(ctx context.Context) {
	if err := s.rdb.FlushAll(ctx); err != nil {
		s.logger.Error("error on flash all from db")
	}

	s.logger.Info("success flush all")
}

func (s *StorageRedis) FlushLimitByBucket(ctx context.Context, bucket string) {
	if err := s.limiter.Reset(ctx, bucket); err != nil {
		s.logger.Errorf("error on reset rate limit by bucket %s: %v", bucket, err)
	}

	s.logger.Infof("success flush rate limit by bucket %s", bucket)
}

func (s *StorageRedis) StoreLimitByLogin(ctx context.Context, login string) (*redis_rate.Result, error) {
	res, err := s.limiter.Allow(ctx, login, redis_rate.PerMinute(s.config.RateLimit.LoginPerMinute+1))
	if err != nil {
		return nil, fmt.Errorf("error on set rate limit by login: %w", err)
	}

	return res, nil
}

func (s *StorageRedis) StoreLimitByPassword(ctx context.Context, password string) (*redis_rate.Result, error) {
	res, err := s.limiter.Allow(ctx, password, redis_rate.PerMinute(s.config.RateLimit.PasswordPerMinute+1))
	if err != nil {
		return nil, fmt.Errorf("error on set rate limit by password: %w", err)
	}

	return res, nil
}

func (s *StorageRedis) StoreLimitByIP(ctx context.Context, ip uint32) (*redis_rate.Result, error) {
	res, err := s.limiter.Allow(ctx, strconv.Itoa(int(ip)), redis_rate.PerMinute(s.config.RateLimit.IPPerMinute+1))
	if err != nil {
		return nil, fmt.Errorf("error on set rate limit by ip: %w", err)
	}

	return res, nil
}

func (s *StorageRedis) ExistInList(ctx context.Context, kindOfList storage.KindOfList, ip net.IP) int64 {
	networks, err := s.rdb.SMembers(ctx, string(kindOfList)).Result()
	if err != nil {
		s.logger.Errorf("can't execute sMember method: %v", err)
		return 0
	}

	for _, network := range networks {
		_, ipNet, err := net.ParseCIDR(network)
		if err != nil {
			s.logger.Errorf("can't parse CIDR: %v", err)
		}

		if ipNet.Contains(ip) {
			return 1
		}
	}

	return 0
}

func (s *StorageRedis) RemoveIPFromListByKind(ctx context.Context, kindOfList storage.KindOfList, netAddress string) {
	s.rdb.SRem(ctx, string(kindOfList), netAddress)
}

func (s *StorageRedis) StoreIPToListByKind(ctx context.Context, kindOfList storage.KindOfList, netAddress string) {
	s.rdb.SAdd(ctx, string(kindOfList), netAddress)
}

func (s *StorageRedis) CheckAccessibilityByLists(
	ctx context.Context, ip net.IP,
) (isAllow bool, isNowFoundedInList bool) {
	if s.ExistInList(ctx, WhiteListKey, ip) == 1 {
		return true, false
	}

	if s.ExistInList(ctx, BlackListKey, ip) == 1 {
		return false, false
	}

	return false, true
}
