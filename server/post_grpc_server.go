package server

import (
	"context"
	"fmt"
	service2 "grpc_identity/internal/service"
	"grpc_identity/pb/v1beta1/post"
	"grpc_identity/server/interceptor"
	"strconv"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type PostGRPCServiceServer struct {
	postService service2.IPostService
	userService service2.IUserService
	post.UnimplementedPostServer
}

func RegisterPostService(postService service2.IPostService, userService service2.IUserService, svr *grpc.Server) {
	post.RegisterPostServer(svr, &PostGRPCServiceServer{
		postService: postService,
		userService: userService,
	})
}

func (p *PostGRPCServiceServer) CreatePost(ctx context.Context, req *post.CreatePostRequest) (*post.CreatePostResponse, error) {
	jwtToken, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		logrus.Errorf("failed to get jwt token: %v", err)
		return nil, fmt.Errorf("failed to get jwt token: %v", err)
	}

	tokenClaimsID, err := interceptor.ExtractTokenFromMetadata(jwtToken)
	if err != nil {
		logrus.Errorf("failed to extract token from metadata: %v", err)
		return nil, fmt.Errorf("failed to extract token from metadata: %v", err)
	}

	tokenID, _ := strconv.Atoi(tokenClaimsID)
	fmt.Println(tokenID)
	userByID, err := p.userService.GetUserByID(ctx, tokenID)
	if err != nil {
		logrus.Errorf("failed to get user: %v", err)
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	if userByID.ID != tokenID {
		logrus.Errorf("unauthorized")
		return nil, fmt.Errorf("unauthorzied")
	}

	createPost, err := p.postService.CreatePost(ctx, req.Title, req.Content, userByID)
	if err != nil {
		logrus.Errorf("failed to create post: %v", err)
		return nil, fmt.Errorf("failed to create post: %v", err)
	}

	postRes := &post.PostMessage{
		Id:      int64(createPost.ID),
		Title:   createPost.Title,
		Content: createPost.Content,
	}

	return &post.CreatePostResponse{Post: postRes}, nil
}

func (p *PostGRPCServiceServer) GetPost(ctx context.Context, req *post.GetPostByIDRequest) (*post.GetPostByIDResponse, error) {
	postByID, err := p.postService.GetPostByID(ctx, int(req.Id))
	if err != nil {
		logrus.Errorf("failed to get post by ID: %v", err)
		return nil, fmt.Errorf("failed to get post by ID: %v", err)
	}

	postRes := &post.PostMessage{
		Id:      int64(postByID.ID),
		Title:   postByID.Title,
		Content: postByID.Content,
	}

	return &post.GetPostByIDResponse{Post: postRes}, nil
}

func (p *PostGRPCServiceServer) GetPostByUser(ctx context.Context, req *post.GetPostByUserRequest) (*post.GetPostByUserResponse, error) {
	jwtToken, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		logrus.Errorf("failed to get jwt token: %v", err)
		return nil, fmt.Errorf("failed to get jwt token: %v", err)
	}

	tokenClaimsID, err := interceptor.ExtractTokenFromMetadata(jwtToken)
	if err != nil {
		logrus.Errorf("failed to extract token from metadata: %v", err)
		return nil, fmt.Errorf("failed to extract token from metadata: %v", err)
	}

	tokenID, _ := strconv.Atoi(tokenClaimsID)
	userByID, err := p.userService.GetUserByID(ctx, tokenID)
	if err != nil {
		logrus.Errorf("failed to get user: %v", err)
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	if userByID.ID != tokenID {
		logrus.Errorf("unauthorized")
		return nil, fmt.Errorf("unauthorzied")
	}

	postsByUserID, err := p.postService.GetPostByUserID(ctx, int(req.UserId))
	if err != nil {
		logrus.Errorf("failed to get post by user: %v", err)
		return nil, fmt.Errorf("failed to get post by user: %v", err)
	}

	var postResponses []*post.PostMessage
	for _, p := range postsByUserID {
		postRes := &post.PostMessage{
			Id:      int64(p.ID),
			Title:   p.Title,
			Content: p.Content,
		}
		postResponses = append(postResponses, postRes)
	}

	return &post.GetPostByUserResponse{Post: postResponses}, nil
}

func (p *PostGRPCServiceServer) DeletePost(ctx context.Context, req *post.DeletePostRequest) (*post.DeletePostResponse, error) {
	jwtToken, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		logrus.Errorf("failed to get jwt token: %v", err)
		return nil, fmt.Errorf("failed to get jwt token: %v", err)
	}

	tokenClaimsID, err := interceptor.ExtractTokenFromMetadata(jwtToken)
	if err != nil {
		logrus.Errorf("failed to extract token from metadata: %v", err)
		return nil, fmt.Errorf("failed to extract token from metadata: %v", err)
	}

	tokenID, _ := strconv.Atoi(tokenClaimsID)
	userByID, err := p.userService.GetUserByID(ctx, tokenID)
	if err != nil {
		logrus.Errorf("failed to get user by ID: %v", err)
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}

	if userByID.ID != tokenID {
		logrus.Errorf("unauthorized")
		return nil, fmt.Errorf("unauthorzied")
	}

	err = p.postService.DeleteByID(ctx, int(req.Id))
	if err != nil {
		logrus.Errorf("failed to delete by ID: %v", err)
		return nil, fmt.Errorf("failed to delete by ID: %v", err)
	}

	return &post.DeletePostResponse{}, nil
}

func (p *PostGRPCServiceServer) UpdatePost(ctx context.Context, req *post.UpdatePostRequest) (*post.UpdatePostResponse, error) {
	jwtToken, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		logrus.Errorf("failed to get jwt token: %v", err)
		return nil, fmt.Errorf("failed to get jwt token: %v", err)
	}

	tokenClaimsID, err := interceptor.ExtractTokenFromMetadata(jwtToken)
	if err != nil {
		logrus.Errorf("failed to extract token from metadata: %v", err)
		return nil, fmt.Errorf("failed to extract token from metadata: %v", err)
	}

	tokenID, _ := strconv.Atoi(tokenClaimsID)
	userByID, err := p.userService.GetUserByID(ctx, tokenID)
	if err != nil {
		logrus.Errorf("failed to get user: %v", err)
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	if userByID.ID != tokenID {
		logrus.Error("unauthorized")
		return nil, fmt.Errorf("unauthorzied")
	}

	updatePost, err := p.postService.UpdatePost(ctx, req.Content, req.Title, int(req.Id))
	if err != nil {
		logrus.Errorf("failed to update post: %v", err)
		return nil, fmt.Errorf("failed to update post: %v", err)
	}

	postMsg := &post.PostMessage{
		Id:      int64(updatePost.ID),
		Title:   updatePost.Title,
		Content: updatePost.Content,
	}

	return &post.UpdatePostResponse{Post: postMsg}, nil
}
