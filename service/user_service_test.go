package service_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"grpc_identity/dto"
	"grpc_identity/ent"
	"grpc_identity/mocks"
	"grpc_identity/service"
	"testing"
	"time"
)

type mockUserDeps struct {
	mockUserRepo *mocks.IUserRepository
}

func createMockUserService() (service.IUserService, *mockUserDeps) {
	mockUserRepo := new(mocks.IUserRepository)
	mockUserService := service.NewUserService(mockUserRepo)
	return mockUserService, &mockUserDeps{mockUserRepo: mockUserRepo}
}

func TestUserService_CreateUser(t *testing.T) {
	t.Run("create user", func(t *testing.T) {
		mockUser := &ent.User{
			ID:       1,
			Name:     "inchul",
			Email:    "inchul@example.com",
			Password: "password",
		}

		mockUserResp := dto.NewUserResponse(mockUser)

		mockUserService, deps := createMockUserService()
		deps.mockUserRepo.On("CreateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(mockUser, nil).Once()

		res, err := mockUserService.CreateUser(context.TODO(), mock.Anything, mock.Anything, mock.Anything)

		require.NoError(t, err)
		require.Equal(t, res, mockUserResp)
		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_GetUserByID(t *testing.T) {
	t.Run("return user by id", func(t *testing.T) {
		mockUser := &ent.User{
			ID:       1,
			Name:     "inchul",
			Email:    "inchul@example.com",
			Password: "password",
		}

		mockUserResp := dto.NewUserResponse(mockUser)

		mockUserService, deps := createMockUserService()
		deps.mockUserRepo.On("GetUserByID", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		res, err := mockUserService.GetUserByID(context.TODO(), 1)

		require.NoError(t, err)
		require.Equal(t, res, mockUserResp)
		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_GetUserByName(t *testing.T) {
	t.Run("return user by name", func(t *testing.T) {
		mockUser := &ent.User{
			ID:        1,
			Name:      "",
			Email:     "",
			Password:  "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockUserResponse := dto.NewUserResponse(mockUser)

		mockUserService, deps := createMockUserService()
		deps.mockUserRepo.On("GetUserByName", mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		res, err := mockUserService.GetUserByName(context.TODO(), mock.Anything)

		require.NoError(t, err)
		require.Equal(t, res, mockUserResponse)
		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_DeleteByID(t *testing.T) {
	t.Run("delete user by id", func(t *testing.T) {
		mockUserService, deps := createMockUserService()
		deps.mockUserRepo.On("DeleteByID", mock.Anything, mock.Anything).Return(nil).Once()
		err := mockUserService.DeleteByID(context.TODO(), 1)

		require.NoError(t, err)
		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	t.Run("update user name", func(t *testing.T) {
		mockUser := &ent.User{
			ID:       1,
			Name:     "inchul",
			Email:    "inchul@example.com",
			Password: "password",
		}

		mockUserResp := dto.NewUserResponse(mockUser)

		mockUserService, deps := createMockUserService()
		deps.mockUserRepo.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockUser, nil).Once()
		res, err := mockUserService.UpdateUser(context.TODO(), mock.Anything, mock.Anything, 1)

		require.NoError(t, err)
		require.Equal(t, res, mockUserResp)

		deps.mockUserRepo.AssertExpectations(t)
	})
}
