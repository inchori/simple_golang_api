package server

import (
	"context"
	"google.golang.org/grpc"
	"grpc_identity/pb/v1beta1/post"
	"grpc_identity/service"
)

type PostGRPCServiceServer struct {
	postService service.IPostService
	userService service.IUserService
	post.UnimplementedPostServer
}

func RegisterPostService(postService service.IPostService, userService service.IUserService, svr *grpc.Server) {
	post.RegisterPostServer(svr, &PostGRPCServiceServer{
		postService: postService,
		userService: userService,
	})
}

func (p *PostGRPCServiceServer) CreatePost(ctx context.Context, req *post.CreatePostRequest) (*post.CreatePostResponse, error) {
	userByID, err := p.userService.GetUserByID(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}

	createPost, err := p.postService.CreatePost(ctx, req.Title, req.Content, userByID)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	postRes := &post.PostMessage{
		Id:      int64(postByID.ID),
		Title:   postByID.Title,
		Content: postByID.Content,
	}

	return &post.GetPostByIDResponse{Post: postRes}, nil
}

func (p *PostGRPCServiceServer) GetPostByUser(ctx context.Context, req *post.GetPostByUserRequest) (*post.GetPostByUserResponse, error) {
	postsByUserID, err := p.postService.GetPostByUserID(ctx, int(req.UserId))
	if err != nil {
		return nil, err
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
	err := p.postService.DeleteByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &post.DeletePostResponse{}, nil
}

func (p *PostGRPCServiceServer) UpdatePost(ctx context.Context, req *post.UpdatePostRequest) (*post.UpdatePostResponse, error) {
	updatePost, err := p.postService.UpdatePost(ctx, req.Content, req.Title, int(req.Id))
	if err != nil {
		return nil, err
	}

	postMsg := &post.PostMessage{
		Id:      int64(updatePost.ID),
		Title:   updatePost.Title,
		Content: updatePost.Content,
	}

	return &post.UpdatePostResponse{Post: postMsg}, nil
}
