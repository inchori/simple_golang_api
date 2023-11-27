package handler

import (
	"context"
	"grpc_identity/dto"
	"grpc_identity/middleware"
	"grpc_identity/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewPostHandler(app fiber.Router, ctx context.Context, postService service.IPostService,
	userService service.IUserService, protected fiber.Handler) {
	app.Post("", protected, CreatePost(ctx, postService, userService))
	app.Get("/:id", GetPostByID(ctx, postService))
	app.Get("/user/:userId", protected, GetPostByUserID(ctx, postService, userService))
	app.Delete("/:id", protected, DeleteByID(ctx, postService, userService))
	app.Put("/:id", protected, UpdatePost(ctx, postService, userService))
}

// CreatePost create a Post
//
//	@Summary	Create a Post
//	@Tags		Post
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.PostCreateRequest	true	"Create a Post Request Body"
//	@Success	200		{object}	dto.PostResponse
//	@Router		/v1/posts [post]
func CreatePost(ctx context.Context, postService service.IPostService, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postCreateRequest := new(dto.PostCreateRequest)
		if err := c.BodyParser(&postCreateRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			tokenClaimsID, err := middleware.ExtractTokenMetadata(c)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			tokenID, _ := strconv.Atoi(tokenClaimsID)
			userByID, err := userService.GetUserByID(ctx, tokenID)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			if strconv.Itoa(userByID.ID) != tokenClaimsID {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "unauthorized",
				})
			}

			post, err := postService.CreatePost(ctx, postCreateRequest.Title, postCreateRequest.Content, userByID)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			} else {
				return c.JSON(dto.NewPostResponse(post))
			}
		}
	}
}

// GetPostByID get a post by ID
//
//	@Summary	Get a Post By ID
//	@Tags		Post
//	@Produce	json
//	@Param		id	path		int	true	"Post ID"
//	@Success	200	{object}	dto.PostResponse
//	@Router		/v1/posts/{id} [get]
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
				return c.JSON(dto.NewPostResponse(postByID))
			}
		}
	}
}

// GetPostByUserID get posts by User ID
//
//	@Summary	Get Posts By User ID
//	@Tags		Post
//	@Produce	json
//	@Param		userId	path		int	true	"User ID"
//	@Success	200		{object}	dto.PostsResponse
//	@Router		/v1/posts/users/{userId} [get]
//	@Security	Bearer
func GetPostByUserID(ctx context.Context, postService service.IPostService, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, err := c.ParamsInt("userId")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			userByID, err := userService.GetUserByID(ctx, userId)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			tokenClaimsID, err := middleware.ExtractTokenMetadata(c)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			if strconv.Itoa(userByID.ID) != tokenClaimsID {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "unauthorized",
				})
			}

			postsByUserID, err := postService.GetPostByUserID(ctx, userId)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			} else {
				var postsResponse []dto.PostResponse
				for _, p := range postsByUserID {
					postsResponse = append(postsResponse, dto.NewPostResponse(p))
				}
				return c.JSON(dto.NewPostsResponse(postsResponse))
			}
		}
	}
}

// DeleteByID delete post by ID
//
//	@Summary	Delete a post
//	@Tags		Post
//	@Produce	json
//	@Param		id	path		int	true	"Post ID"
//	@Success	200	{object}	string
//	@Router		/v1/posts/{id} [delete]
//	@Security	Bearer
func DeleteByID(ctx context.Context, postService service.IPostService, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			tokenClaimsID, err := middleware.ExtractTokenMetadata(c)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			userID, _ := strconv.Atoi(tokenClaimsID)
			userByID, err := userService.GetUserByID(ctx, userID)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			if strconv.Itoa(userByID.ID) != tokenClaimsID {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "unauthorized",
				})
			}

			err = postService.DeleteByID(ctx, id)
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

// UpdatePost update post
//
//	@Summary	Update Post
//	@Tags		Post
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.PostUpdateRequest	true	"Update Post Request Body"
//	@Success	200		{object}	dto.PostResponse
//	@Router		/v1/posts/{id} [put]
//
//	@Security	Bearer
func UpdatePost(ctx context.Context, postService service.IPostService, userService service.IUserService) fiber.Handler {
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
			tokenClaimsID, err := middleware.ExtractTokenMetadata(c)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			userID, _ := strconv.Atoi(tokenClaimsID)
			userByID, err := userService.GetUserByID(ctx, userID)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			if strconv.Itoa(userByID.ID) != tokenClaimsID {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "unauthorized",
				})
			}

			post, err := postService.UpdatePost(ctx, postUpdateRequest.Title, postUpdateRequest.Content, id)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			} else {
				return c.JSON(dto.NewPostResponse(post))
			}
		}
	}
}
