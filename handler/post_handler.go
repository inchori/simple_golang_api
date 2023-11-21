package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"grpc_identity/service"
)

func NewPostHandler(app fiber.Router, ctx context.Context, postService service.IPostService) {
}
