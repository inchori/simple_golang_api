// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"
	ent "grpc_identity/ent"

	mock "github.com/stretchr/testify/mock"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, name, email, password
func (_m *IUserRepository) CreateUser(ctx context.Context, name string, email string, password string) (*ent.User, error) {
	ret := _m.Called(ctx, name, email, password)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*ent.User, error)); ok {
		return rf(ctx, name, email, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *ent.User); ok {
		r0 = rf(ctx, name, email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *IUserRepository) DeleteByID(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *IUserRepository) GetUserByID(ctx context.Context, id int) (*ent.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*ent.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *ent.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByName provides a mock function with given fields: ctx, name
func (_m *IUserRepository) GetUserByName(ctx context.Context, name string) (*ent.User, error) {
	ret := _m.Called(ctx, name)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*ent.User, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *ent.User); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, name, password, id
func (_m *IUserRepository) UpdateUser(ctx context.Context, name string, password string, id int) (*ent.User, error) {
	ret := _m.Called(ctx, name, password, id)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) (*ent.User, error)); ok {
		return rf(ctx, name, password, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) *ent.User); ok {
		r0 = rf(ctx, name, password, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int) error); ok {
		r1 = rf(ctx, name, password, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIUserRepository creates a new instance of IUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserRepository {
	mock := &IUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
