package lickybucket

import (
	"context"
	"sync"

	"github.com/Tyrqvir/anti-brute-force/internal/limiter"
	"github.com/Tyrqvir/anti-brute-force/internal/logger"
	"github.com/Tyrqvir/anti-brute-force/internal/storage"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
)

var (
	mutex sync.Mutex
	wg    sync.WaitGroup
)

type Limiter struct {
	limiter.Algorithm

	storage storage.Storage
	logger  logger.Logger
}

func New(storage storage.Storage, logger logger.Logger) *Limiter {
	return &Limiter{
		storage: storage,
		logger:  logger,
	}
}

func (l *Limiter) IsLimit(ctx context.Context, request *api.AuthorisationRequest) bool {
	isLimit := false

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			if isLimit {
				return
			}

			mutex.Lock()
			defer mutex.Unlock()

			switch i {
			case 0:
				result := l.LimitByLogin(ctx, request)

				if result {
					isLimit = result
				}
			case 1:
				result := l.LimitByPassword(ctx, request)

				if result {
					isLimit = result
				}

			case 2:
				result := l.LimitByIP(ctx, request)

				if result {
					isLimit = result
				}
			}
		}(i)
	}

	wg.Wait()

	return isLimit
}

func (l *Limiter) LimitByIP(ctx context.Context, request *api.AuthorisationRequest) bool {
	storeLimitByIP, err := l.storage.StoreLimitByIP(ctx, request.GetIp())
	if err != nil {
		l.logger.Errorf("error on store: %v", err)
	} else {
		l.logger.Infof("limit by ip: allowed %d, remaining %d", storeLimitByIP.Allowed, storeLimitByIP.Remaining)
		if storeLimitByIP.Allowed == 0 {
			return true
		}
	}

	return false
}

func (l *Limiter) LimitByLogin(ctx context.Context, request *api.AuthorisationRequest) bool {
	storeLimitByLogin, err := l.storage.StoreLimitByLogin(ctx, request.GetLogin())
	if err != nil {
		l.logger.Errorf("error on store: %v", err)
	} else {
		l.logger.Infof("limit by login: allowed %d, remaining %d", storeLimitByLogin.Allowed, storeLimitByLogin.Remaining)
		if storeLimitByLogin.Allowed == 0 {
			return true
		}
	}

	return false
}

func (l *Limiter) LimitByPassword(ctx context.Context, request *api.AuthorisationRequest) bool {
	storeLimitByPassword, err := l.storage.StoreLimitByPassword(ctx, request.GetPassword())
	if err != nil {
		l.logger.Errorf("error on store: %v", err)
	} else {
		l.logger.Infof(
			"limit by password: allowed %d, remaining %d",
			storeLimitByPassword.Allowed, storeLimitByPassword.Remaining,
		)
		if storeLimitByPassword.Allowed == 0 {
			return true
		}
	}
	return false
}

func (l *Limiter) ClearBucket(ctx context.Context, request *api.ResetBucketRequest) {
	l.storage.FlushLimitByBucket(ctx, request.GetIp())
	l.storage.FlushLimitByBucket(ctx, request.GetLogin())
}
