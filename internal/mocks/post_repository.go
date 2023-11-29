// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"

	ent "github.com/inchori/grpc_identity/internal/ent"
	mock "github.com/stretchr/testify/mock"
)

// IPostRepository is an autogenerated mock type for the IPostRepository type
type IPostRepository struct {
	mock.Mock
}

// CreatePost provides a mock function with given fields: ctx, title, content, user
func (_m *IPostRepository) CreatePost(ctx context.Context, title string, content string, user *ent.User) (*ent.Post, error) {
	ret := _m.Called(ctx, title, content, user)

	var r0 *ent.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *ent.User) (*ent.Post, error)); ok {
		return rf(ctx, title, content, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *ent.User) *ent.Post); ok {
		r0 = rf(ctx, title, content, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *ent.User) error); ok {
		r1 = rf(ctx, title, content, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *IPostRepository) DeleteByID(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPostByID provides a mock function with given fields: ctx, id
func (_m *IPostRepository) GetPostByID(ctx context.Context, id int) (*ent.Post, error) {
	ret := _m.Called(ctx, id)

	var r0 *ent.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*ent.Post, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *ent.Post); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostByUserID provides a mock function with given fields: ctx, userID
func (_m *IPostRepository) GetPostByUserID(ctx context.Context, userID int) ([]*ent.Post, error) {
	ret := _m.Called(ctx, userID)

	var r0 []*ent.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]*ent.Post, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []*ent.Post); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePost provides a mock function with given fields: ctx, title, content, id
func (_m *IPostRepository) UpdatePost(ctx context.Context, title string, content string, id int) (*ent.Post, error) {
	ret := _m.Called(ctx, title, content, id)

	var r0 *ent.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) (*ent.Post, error)); ok {
		return rf(ctx, title, content, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) *ent.Post); ok {
		r0 = rf(ctx, title, content, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, int) error); ok {
		r1 = rf(ctx, title, content, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIPostRepository creates a new instance of IPostRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPostRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IPostRepository {
	mock := &IPostRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}