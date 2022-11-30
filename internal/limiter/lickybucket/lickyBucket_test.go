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

func dummyLimiter(storage *mocks.Storage) *Limiter {
	logger := appLogger.New("info")
	return &Limiter{
		storage: storage,
		logger:  logger,
	}
}

func dummyRequest() *api.AccessCheckRequest {
	return &api.AccessCheckRequest{
		Login:    "login",
		Password: "password",
		Ip:       3232235777, // "192.168.1.1"
	}
}

func TestLimiter_limitByLogin(t *testing.T) { //nolint:dupl
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		request *api.AccessCheckRequest
	}
	tests := []struct {
		name    string
		prepare func(s *mocks.Storage)
		args    args
		want    bool
	}{
		{
			name: "normal access check by login",
			prepare: func(s *mocks.Storage) {
				s.On("StoreLimitByLogin", ctx, "login").Return(&redis_rate.Result{
					Allowed: 1,
				}, nil)
			},
			args: args{
				ctx:     ctx,
				request: dummyRequest(),
			},
			want: false,
		},
		{
			name: "ddos access check by login",
			prepare: func(s *mocks.Storage) {
				s.On("StoreLimitByLogin", ctx, "login").Return(&redis_rate.Result{
					Allowed: 0,
				}, nil)
			},
			args: args{
				ctx:     ctx,
				request: dummyRequest(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &mocks.Storage{}
			tt.prepare(store)
			l := dummyLimiter(store)
			assert.Equalf(
				t,
				tt.want,
				l.LimitByLogin(tt.args.ctx, tt.args.request),
				"LimitByLogin(%v, %v)",
				tt.args.ctx,
				tt.args.request,
			)
		})
	}
}

func TestLimiter_limitByPassword(t *testing.T) { //nolint:dupl
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		request *api.AccessCheckRequest
	}
	tests := []struct {
		name    string
		prepare func(s *mocks.Storage)
		args    args
		want    bool
	}{
		{
			name: "normal access check by password",
			prepare: func(s *mocks.Storage) {
				s.On("StoreLimitByPassword", ctx, "password").Return(&redis_rate.Result{
					Allowed: 1,
				}, nil)
			},
			args: args{
				ctx:     ctx,
				request: dummyRequest(),
			},
			want: false,
		},
		{
			name: "ddos access check by password",
			prepare: func(s *mocks.Storage) {
				s.On("StoreLimitByPassword", ctx, "password").Return(&redis_rate.Result{
					Allowed: 0,
				}, nil)
			},
			args: args{
				ctx:     ctx,
				request: dummyRequest(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &mocks.Storage{}
			tt.prepare(store)
			l := dummyLimiter(store)
			assert.Equalf(
				t,
				tt.want,
				l.LimitByPassword(tt.args.ctx, tt.args.request),
				"LimitByPassword(%v, %v)",
				tt.args.ctx,
				tt.args.request,
			)
		})
	}
}

func TestLimiter_limitByIP(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx     context.Context
		request *api.AccessCheckRequest
	}
	tests := []struct {
		name    string
		prepare func(s *mocks.Storage)
		args    args
		want    bool
	}{
		{
			name: "normal access check by ip",
			prepare: func(s *mocks.Storage) {
				var ip uint32 = 3232235777
				s.On("StoreLimitByIP", ctx, ip).Return(&redis_rate.Result{
					Allowed: 1,
				}, nil)
			},
			args: args{
				ctx:     ctx,
				request: dummyRequest(),
			},
			want: false,
		},
		{
			name: "ddos access check by ip",
			prepare: func(s *mocks.Storage) {
				var ip uint32 = 3232235777
				s.On("StoreLimitByIP", ctx, ip).Return(&redis_rate.Result{
					Allowed: 0,
				}, nil)
			},
			args: args{
				ctx:     ctx,
				request: dummyRequest(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &mocks.Storage{}
			tt.prepare(store)
			l := dummyLimiter(store)
			assert.Equalf(
				t,
				tt.want,
				l.LimitByIP(tt.args.ctx, tt.args.request),
				"LimitByIP(%v, %v)", tt.args.ctx, tt.args.request,
			)
		})
	}
}

func TestLimiter_IsLimit(t *testing.T) {
	ctx := context.Background()

	store := &mocks.Storage{}

	l := dummyLimiter(store)

	request := dummyRequest()

	t.Run("test authorization without block", func(t *testing.T) {
		result := &redis_rate.Result{
			Allowed: 1,
		}
		var ip uint32 = 3232235777
		store.On("StoreLimitByIP", ctx, ip).Return(result, nil)
		store.On("StoreLimitByLogin", ctx, "login").Return(result, nil)
		store.On("StoreLimitByPassword", ctx, "password").Return(result, nil)

		isLimit := l.IsLimit(ctx, request)

		assert.False(t, isLimit)
	})
}
