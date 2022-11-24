package usercases

import (
	"context"
	"net"

	"github.com/Tyrqvir/anti-brute-force/internal/limiter/lickybucket"
	"github.com/Tyrqvir/anti-brute-force/internal/storage/redis"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AntiBruteForceCase struct {
	api.UnimplementedAntiBruteForceServiceServer

	limiter *lickybucket.Limiter
	storage *redis.StorageRedis
}

//go:generate mockery --name AntiBruteForceServiceServer --dir ./ --output ./../../internal/mocks
type AntiBruteForceServiceServer interface {
	AccessCheck(context.Context, *api.AccessCheckRequest) (*api.AccessCheckResponse, error)
	ResetBucket(context.Context, *api.ResetBucketRequest) (*emptypb.Empty, error)
	AddToBlackList(context.Context, *api.ListCases) (*emptypb.Empty, error)
	RemoveFromBlackList(context.Context, *api.ListCases) (*emptypb.Empty, error)
	AddToWhiteList(context.Context, *api.ListCases) (*emptypb.Empty, error)
	RemoveFromWhiteList(context.Context, *api.ListCases) (*emptypb.Empty, error)
	ExistInWhiteList(context.Context, *api.ListCases) (*api.ExistInListResponse, error)
	ExistInBlackList(context.Context, *api.ListCases) (*api.ExistInListResponse, error)
	mustEmbedUnimplementedAntiBruteForceServiceServer()
}

func NewUseCase(limiter *lickybucket.Limiter, storage *redis.StorageRedis) *AntiBruteForceCase {
	return &AntiBruteForceCase{
		limiter: limiter,
		storage: storage,
	}
}

func (u AntiBruteForceCase) Authorisation(
	ctx context.Context,
	request *api.AccessCheckRequest,
) (*api.AccessCheckResponse, error) {
	ip := net.ParseIP(request.GetIp())

	accessibility, notExistInLists := u.storage.CheckAccessibilityByLists(ctx, ip)

	if notExistInLists {
		isLimit := u.limiter.IsLimit(ctx, request)
		return &api.AccessCheckResponse{
			Ok: !isLimit,
		}, nil
	}

	return &api.AccessCheckResponse{
		Ok: accessibility,
	}, nil
}

func (u AntiBruteForceCase) ResetBucket(ctx context.Context, request *api.ResetBucketRequest) (*emptypb.Empty, error) {
	u.limiter.ClearBucket(ctx, request)

	return &emptypb.Empty{}, nil
}

func (u AntiBruteForceCase) AddToBlackList(ctx context.Context, cases *api.ListCases) (*emptypb.Empty, error) {
	ipNet, err := getIPNet(cases.GetSubNetwork())
	if err != nil {
		return nil, err
	}

	u.storage.StoreIPToListByKind(ctx, redis.BlackListKey, ipNet.String())

	return &emptypb.Empty{}, nil
}

func (u AntiBruteForceCase) RemoveFromBlackList(ctx context.Context, cases *api.ListCases) (*emptypb.Empty, error) {
	ipNet, err := getIPNet(cases.GetSubNetwork())
	if err != nil {
		return nil, err
	}

	u.storage.RemoveIPFromListByKind(ctx, redis.BlackListKey, ipNet.String())

	return &emptypb.Empty{}, nil
}

func (u AntiBruteForceCase) AddToWhiteList(ctx context.Context, cases *api.ListCases) (*emptypb.Empty, error) {
	ipNet, err := getIPNet(cases.GetSubNetwork())
	if err != nil {
		return nil, err
	}

	u.storage.StoreIPToListByKind(ctx, redis.WhiteListKey, ipNet.String())

	return &emptypb.Empty{}, nil
}

func (u AntiBruteForceCase) ExistInWhiteList(
	ctx context.Context,
	cases *api.ListCases,
) (*api.ExistInListResponse, error) {
	ipNet, err := getIPNet(cases.GetSubNetwork())
	if err != nil {
		return nil, err
	}
	exist := u.storage.ExistInList(ctx, redis.WhiteListKey, ipNet.IP)

	return &api.ExistInListResponse{
		Ok: exist,
	}, nil
}

func (u AntiBruteForceCase) ExistInBlackList(
	ctx context.Context,
	cases *api.ListCases,
) (*api.ExistInListResponse, error) {
	ipNet, err := getIPNet(cases.GetSubNetwork())
	if err != nil {
		return nil, err
	}
	exist := u.storage.ExistInList(ctx, redis.BlackListKey, ipNet.IP)
	return &api.ExistInListResponse{
		Ok: exist,
	}, nil
}

func (u AntiBruteForceCase) RemoveFromWhiteList(ctx context.Context, cases *api.ListCases) (*emptypb.Empty, error) {
	ipNet, err := getIPNet(cases.GetSubNetwork())
	if err != nil {
		return nil, err
	}

	u.storage.RemoveIPFromListByKind(ctx, redis.WhiteListKey, ipNet.String())

	return &emptypb.Empty{}, nil
}

func getIPNet(network string) (*net.IPNet, error) {
	_, ipNet, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}

	return ipNet, nil
}
