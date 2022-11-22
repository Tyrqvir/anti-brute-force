// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/Tyrqvir/anti-brute-force/proto/api"

	mock "github.com/stretchr/testify/mock"
)

// Algorithm is an autogenerated mock type for the Algorithm type
type Algorithm struct {
	mock.Mock
}

// ClearBucket provides a mock function with given fields: ctx, request
func (_m *Algorithm) ClearBucket(ctx context.Context, request *api.ResetBucketRequest) {
	_m.Called(ctx, request)
}

// IsLimit provides a mock function with given fields: ctx, request
func (_m *Algorithm) IsLimit(ctx context.Context, request *api.AuthorisationRequest) bool {
	ret := _m.Called(ctx, request)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *api.AuthorisationRequest) bool); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LimitByIP provides a mock function with given fields: ctx, request
func (_m *Algorithm) LimitByIP(ctx context.Context, request *api.AuthorisationRequest) bool {
	ret := _m.Called(ctx, request)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *api.AuthorisationRequest) bool); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LimitByLogin provides a mock function with given fields: ctx, request
func (_m *Algorithm) LimitByLogin(ctx context.Context, request *api.AuthorisationRequest) bool {
	ret := _m.Called(ctx, request)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *api.AuthorisationRequest) bool); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LimitByPassword provides a mock function with given fields: ctx, request
func (_m *Algorithm) LimitByPassword(ctx context.Context, request *api.AuthorisationRequest) bool {
	ret := _m.Called(ctx, request)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *api.AuthorisationRequest) bool); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewAlgorithm interface {
	mock.TestingT
	Cleanup(func())
}

// NewAlgorithm creates a new instance of Algorithm. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAlgorithm(t mockConstructorTestingTNewAlgorithm) *Algorithm {
	mock := &Algorithm{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
