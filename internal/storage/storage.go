package storage

import (
	"context"
	"net"

	"github.com/go-redis/redis_rate/v9"
)

type KindOfList string

//go:generate mockery --name Storage --dir ./ --output ./../../internal/mocks
type Storage interface {
	FlushLimitByBucket(ctx context.Context, bucket string)
	StoreLimitByLogin(ctx context.Context, login string) (*redis_rate.Result, error)
	StoreLimitByPassword(ctx context.Context, password string) (*redis_rate.Result, error)
	StoreLimitByIP(ctx context.Context, ip uint32) (*redis_rate.Result, error)
	ExistInList(ctx context.Context, kindOfList KindOfList, ip net.IP) int64
	RemoveIPFromListByKind(ctx context.Context, kindOfList KindOfList, netAddress string)
	StoreIPToListByKind(ctx context.Context, kindOfList KindOfList, netAddress string)
	CheckAccessibilityByLists(ctx context.Context, ip net.IP) (isAllow bool, isNowFoundedInList bool)
}
