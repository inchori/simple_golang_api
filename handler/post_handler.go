package handler

import (
	"context"
	"grpc_identity/dto"
	"grpc_identity/middleware"
	"grpc_identity/service"

	"github.com/gofiber/fiber/v2"
)

func NewPostHandler(app fiber.Router, ctx context.Context, postService service.IPostService, protected fiber.Handler) {
	app.Post("", middleware.Protected(), CreatePost(ctx, postService))
	app.Get("/:id", GetPostByID(ctx, postService))
	app.Delete("/:id", protected, DeleteByID(ctx, postService))
	app.Put("/:id", protected, UpdatePost(ctx, postService))
}

func CreatePost(ctx context.Context, postService service.IPostService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postCreateRequest := new(dto.PostCreateRequest)
		if err := c.BodyParser(&postCreateRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			post, err := postService.CreatePost(ctx, postCreateRequest.Title, postCreateRequest.Content)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			} else {
				return c.JSON(post)
			}
		}
	}
}

func GetPostByID(ctx context.Context, postService service.IPostService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			postByID, err := postService.GetPostByID(ctx, id)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			} else {
				return c.JSON(postByID)
			}
		}
	}
}

func DeleteByID(ctx context.Context, postService service.IPostService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			err := postService.DeleteByID(ctx, id)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			} else {
				return c.JSON(fiber.Map{
					"message": "successfully delete post",
				})
			}
		}
	}
}

func UpdatePost(ctx context.Context, postService service.IPostService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		postUpdateRequest := new(dto.PostUpdateRequest)
		if err := c.BodyParser(&postUpdateRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			postResponse, err := postService.UpdatePost(ctx, postUpdateRequest.Title, postUpdateRequest.Content, id)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			} else {
				return c.JSON(postResponse)
			}
		}
	}
}
