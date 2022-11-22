package lickybucket

import (
	"context"
	"testing"

	appLogger "github.com/Tyrqvir/anti-brute-force/internal/logger"
	"github.com/Tyrqvir/anti-brute-force/internal/mocks"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
	"github.com/go-redis/redis_rate/v9"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func dummyLimiter(storage *mocks.Storage) *Limiter {
	logger := appLogger.New("info")
	return &Limiter{
		storage: storage,
		logger:  logger,
	}
}

func dummyRequest() *api.AuthorisationRequest {
	return &api.AuthorisationRequest{
		Login:    "login",
		Password: "password",
		Ip:       "192.168.1.1",
	}
}

func TestLimiter_limitByLogin(t *testing.T) {
	t.Run("normal authorization by login", func(t *testing.T) {
		storage := &mocks.Storage{}
		storage.On("StoreLimitByLogin", ctx, "login").Return(&redis_rate.Result{
			Allowed: 1,
		}, nil)

		l := dummyLimiter(storage)

		isLimit := l.LimitByLogin(ctx, dummyRequest())

		assert.False(t, isLimit)
	})

	t.Run("ddos authorization by login", func(t *testing.T) {
		storage := &mocks.Storage{}
		storage.On("StoreLimitByLogin", ctx, "login").Return(&redis_rate.Result{
			Allowed: 0,
		}, nil)

		l := dummyLimiter(storage)

		isLimit := l.LimitByLogin(ctx, dummyRequest())

		assert.True(t, isLimit)
	})
}

func TestLimiter_limitByPassword(t *testing.T) {
	t.Run("normal authorization by password", func(t *testing.T) {
		storage := &mocks.Storage{}
		storage.On("StoreLimitByPassword", ctx, "password").Return(&redis_rate.Result{
			Allowed: 1,
		}, nil)

		l := dummyLimiter(storage)

		isLimit := l.LimitByPassword(ctx, dummyRequest())

		assert.False(t, isLimit)
	})

	t.Run("ddos authorization by password", func(t *testing.T) {
		storage := &mocks.Storage{}
		storage.On("StoreLimitByPassword", ctx, "password").Return(&redis_rate.Result{
			Allowed: 0,
		}, nil)

		l := dummyLimiter(storage)

		isLimit := l.LimitByPassword(ctx, dummyRequest())

		assert.True(t, isLimit)
	})
}

func TestLimiter_limitByIP(t *testing.T) {
	t.Run("normal authorization by ip", func(t *testing.T) {
		storage := &mocks.Storage{}
		storage.On("StoreLimitByIP", ctx, "192.168.1.1").Return(&redis_rate.Result{
			Allowed: 1,
		}, nil)

		l := dummyLimiter(storage)

		isLimit := l.LimitByIP(ctx, dummyRequest())

		assert.False(t, isLimit)
	})

	t.Run("ddos authorization by ip", func(t *testing.T) {
		storage := &mocks.Storage{}
		storage.On("StoreLimitByIP", ctx, "192.168.1.1").Return(&redis_rate.Result{
			Allowed: 0,
		}, nil)

		l := dummyLimiter(storage)

		isLimit := l.LimitByIP(ctx, dummyRequest())

		assert.True(t, isLimit)
	})
}

func TestLimiter_IsLimit(t *testing.T) {
	storage := &mocks.Storage{}

	l := dummyLimiter(storage)

	request := dummyRequest()

	t.Run("test authorization without block", func(t *testing.T) {
		result := &redis_rate.Result{
			Allowed: 1,
		}
		storage.On("StoreLimitByIP", ctx, "192.168.1.1").Return(result, nil)
		storage.On("StoreLimitByLogin", ctx, "login").Return(result, nil)
		storage.On("StoreLimitByPassword", ctx, "password").Return(result, nil)

		isLimit := l.IsLimit(ctx, request)

		assert.False(t, isLimit)
	})
}
