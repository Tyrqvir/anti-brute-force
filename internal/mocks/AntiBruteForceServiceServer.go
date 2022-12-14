// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/Tyrqvir/anti-brute-force/proto/api"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	mock "github.com/stretchr/testify/mock"
)

// AntiBruteForceServiceServer is an autogenerated mock type for the AntiBruteForceServiceServer type
type AntiBruteForceServiceServer struct {
	mock.Mock
}

// AccessCheck provides a mock function with given fields: _a0, _a1
func (_m *AntiBruteForceServiceServer) AccessCheck(_a0 context.Context, _a1 *api.AccessCheckRequest) (*api.AccessCheckResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *api.AccessCheckResponse
	if rf, ok := ret.Get(0).(func(context.Context, *api.AccessCheckRequest) *api.AccessCheckResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.AccessCheckResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *api.AccessCheckRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddToBlackList provides a mock function with given fields: _a0, _a1
func (_m *AntiBruteForceServiceServer) AddToBlackList(_a0 context.Context, _a1 *api.ListCases) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *api.ListCases) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *api.ListCases) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddToWhiteList provides a mock function with given fields: _a0, _a1
func (_m *AntiBruteForceServiceServer) AddToWhiteList(_a0 context.Context, _a1 *api.ListCases) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *api.ListCases) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *api.ListCases) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistInBlackList provides a mock function with given fields: _a0, _a1
func (_m *AntiBruteForceServiceServer) ExistInBlackList(_a0 context.Context, _a1 *api.ListCases) (*api.ExistInListResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *api.ExistInListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *api.ListCases) *api.ExistInListResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ExistInListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *api.ListCases) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistInWhiteList provides a mock function with given fields: _a0, _a1
func (_m *AntiBruteForceServiceServer) ExistInWhiteList(_a0 context.Context, _a1 *api.ListCases) (*api.ExistInListResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *api.ExistInListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *api.ListCases) *api.ExistInListResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ExistInListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *api.ListCases) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveFromBlackList provides a mock function with given fields: _a0, _a1
func (_m *AntiBruteForceServiceServer) RemoveFromBlackList(_a0 context.Context, _a1 *api.ListCases) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *api.ListCases) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *api.ListCases) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveFromWhiteList provides a mock function with given fields: _a0, _a1
func (_m *AntiBruteForceServiceServer) RemoveFromWhiteList(_a0 context.Context, _a1 *api.ListCases) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *api.ListCases) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *api.ListCases) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetBucket provides a mock function with given fields: _a0, _a1
func (_m *AntiBruteForceServiceServer) ResetBucket(_a0 context.Context, _a1 *api.ResetBucketRequest) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *emptypb.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *api.ResetBucketRequest) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *api.ResetBucketRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedAntiBruteForceServiceServer provides a mock function with given fields:
func (_m *AntiBruteForceServiceServer) mustEmbedUnimplementedAntiBruteForceServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewAntiBruteForceServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewAntiBruteForceServiceServer creates a new instance of AntiBruteForceServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAntiBruteForceServiceServer(t mockConstructorTestingTNewAntiBruteForceServiceServer) *AntiBruteForceServiceServer {
	mock := &AntiBruteForceServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
