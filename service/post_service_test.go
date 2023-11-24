package service_test

import (
	"context"
	"grpc_identity/ent"
	"grpc_identity/mocks"
	"grpc_identity/service"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockPostDeps struct {
	mockPostRepo *mocks.IPostRepository
}

func createMockPostService() (service.IPostService, *mockPostDeps) {
	mockPostRepo := new(mocks.IPostRepository)
	mockPostService := service.NewPostService(mockPostRepo)
	return mockPostService, &mockPostDeps{mockPostRepo: mockPostRepo}
}

func TestPostService_CreatePost(t *testing.T) {
	t.Run("create post", func(t *testing.T) {
		mockUser := &ent.User{
			ID:    1,
			Email: "inchul@example.com",
			Name:  "inchul",
		}

		mockPost := &ent.Post{
			ID:      1,
			Title:   "title",
			Content: "content",
		}

		mockPostService, deps := createMockPostService()
		deps.mockPostRepo.On("CreatePost", mock.Anything, mock.Anything, mock.Anything, mockUser).Return(mockPost, nil).Once()
		res, err := mockPostService.CreatePost(context.TODO(), mock.Anything, mock.Anything, mockUser)

		require.NoError(t, err)
		require.Equal(t, res, mockPost)

		deps.mockPostRepo.AssertExpectations(t)
	})
}

func TestPostService_GetPostByID(t *testing.T) {
	t.Run("return post by id", func(t *testing.T) {
		mockPost := &ent.Post{
			ID:      1,
			Title:   "title",
			Content: "content",
		}

		mockPostService, deps := createMockPostService()
		deps.mockPostRepo.On("GetPostByID", mock.Anything, mock.Anything).Return(mockPost, nil).Once()
		res, err := mockPostService.GetPostByID(context.TODO(), 1)

		require.NoError(t, err)
		require.Equal(t, res, mockPost)

		deps.mockPostRepo.AssertExpectations(t)
	})
}

func TestPostService_DeleteByID(t *testing.T) {
	t.Run("delete post by id", func(t *testing.T) {
		mockPostService, deps := createMockPostService()
		deps.mockPostRepo.On("DeleteByID", mock.Anything, mock.Anything).Return(nil).Once()
		err := mockPostService.DeleteByID(context.TODO(), 1)

		require.NoError(t, err)

		deps.mockPostRepo.AssertExpectations(t)
	})
}

func TestPostService_GetPostByUserID(t *testing.T) {
	t.Run("return posts by user Id", func(t *testing.T) {
		mockUser := &ent.User{
			ID:       1,
			Email:    "inchul@example.com",
			Password: "example",
		}

		mockPost := &ent.Post{
			ID:      1,
			Title:   "title",
			Content: "content",
		}

		var mockPostResponse []*ent.Post
		mockPosts := append(mockPostResponse, mockPost)

		mockPostService, deps := createMockPostService()
		deps.mockPostRepo.On("GetPostByUserID", mock.Anything, mock.Anything).Return(mockPosts, nil).Once()
		res, err := mockPostService.GetPostByUserID(context.TODO(), mockUser.ID)

		require.NoError(t, err)
		require.Equal(t, res, mockPosts)
	})
}

func TestPostService_UpdatePost(t *testing.T) {
	t.Run("update post title", func(t *testing.T) {
		mockPost := &ent.Post{
			ID:      1,
			Title:   "title",
			Content: "content",
		}

		mockPostService, deps := createMockPostService()
		deps.mockPostRepo.On("UpdatePost", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(mockPost, nil).Once()
		res, err := mockPostService.UpdatePost(context.TODO(), mock.Anything, mock.Anything, 1)

		require.NoError(t, err)
		require.Equal(t, res, mockPost)

		deps.mockPostRepo.AssertExpectations(t)
	})
}
