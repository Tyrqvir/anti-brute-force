package limiter

import (
	"context"

	"github.com/Tyrqvir/anti-brute-force/proto/api"
)

//go:generate mockery --name Algorithm --dir ./ --output ./../../internal/mocks
type Algorithm interface {
	IsLimit(ctx context.Context, request *api.AuthorisationRequest) bool
	LimitByIP(ctx context.Context, request *api.AuthorisationRequest) bool
	LimitByLogin(ctx context.Context, request *api.AuthorisationRequest) bool
	LimitByPassword(ctx context.Context, request *api.AuthorisationRequest) bool
	ClearBucket(ctx context.Context, request *api.ResetBucketRequest)
}
