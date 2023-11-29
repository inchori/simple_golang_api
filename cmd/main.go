package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/inchori/grpc_identity/config"
	_ "github.com/inchori/grpc_identity/docs"
	"github.com/inchori/grpc_identity/internal/database"
	"github.com/inchori/grpc_identity/internal/handler"
	"github.com/inchori/grpc_identity/internal/middleware"
	"github.com/inchori/grpc_identity/internal/repository"
	"github.com/inchori/grpc_identity/internal/service"
	"github.com/inchori/grpc_identity/server"
	"github.com/inchori/grpc_identity/server/interceptor"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

// @title						Simple Blog CRUD API
// @version					1.0
// @termsOfService				http://swagger.io/terms/
// @contact.name				Inchul Song
// @contact.email				sic61695@gmail.com
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token.
func main() {
	app := fiber.New()

	loadConfig, err := config.LoadConfigFile()
	if err != nil {
		logrus.Fatalf("failed to load config file: %v", err)
	}

	dbClient := database.ConnectDB(*loadConfig)
	ctx := context.Background()
	if err := dbClient.Schema.Create(ctx); err != nil {
		logrus.Fatalf("failed to migrate database schema: %v", err)
	}

	userRepository := repository.NewUserRepository(dbClient.User)
	userService := service.NewUserService(userRepository)
	postRepository := repository.NewPostRepository(dbClient.Post)
	postService := service.NewPostService(postRepository)

	if loadConfig.Server == "grpc" {
		lis, err := net.Listen("tcp", ":"+loadConfig.GRPCPort)
		if err != nil {
			logrus.Fatalf("failed to lieten on %s port: %v", loadConfig.GRPCPort, err)
		}

		methods := []string{
			"proto.v1beta1.post.Post/CreatePost",
			"proto.v1beta1.post.Post/UpdatePost",
			"proto.v1beta1.post.Post/DeletePost",
			"proto.v1beta1.user.Post/GetPostByUser",
			"proto.v1beta1.user.User/UpdateUser",
			"proto.v1beta1.user.User/DeleteUser",
		}

		logger := logrus.New()

		jwtInterceptor := interceptor.NewJWTInterceptor(methods)

		grpcSvr := grpc.NewServer(grpc.ChainUnaryInterceptor(
			jwtInterceptor.Interceptor,
			logging.UnaryServerInterceptor(interceptor.LoggerInterceptor(logger)),
		))

		server.RegisterAuthService(userService, grpcSvr)
		server.RegisterUserService(userService, grpcSvr)
		server.RegisterPostService(postService, userService, grpcSvr)
		reflection.Register(grpcSvr)

		go func() {
			if loadConfig.GatewayEnabled == "true" {
				newMux, err := server.GatewayServer(lis.Addr().String())
				if err != nil {
					logrus.Fatalf("failed to serve configure gateway server: %v", err)
				}
				logrus.Infof("gRPC server with gateway is running on %s port", loadConfig.GatewayPort)
				if err := http.ListenAndServe(":"+loadConfig.GatewayPort, newMux); err != nil {
					logrus.Fatalf("failed to listen gateway server: %v", err)
				}
			}
		}()

		logrus.Infof("gRPC server is running on %s port", loadConfig.GRPCPort)
		if err := grpcSvr.Serve(lis); err != nil {
			logrus.Fatalf("failed to serve gRPC server: %v", err)
		}

	} else if loadConfig.Server == "http" {
		protected := middleware.Protected()

		handler.NewLoginHandler(app.Group("/v1/auth"), context.Background(), userService)
		handler.NewUserHandler(app.Group("/v1/users"), context.Background(), userService, protected)
		handler.NewPostHandler(app.Group("/v1/posts"), context.Background(), postService, userService, protected)

		app.Get("/swagger/*", swagger.HandlerDefault)
		logrus.Fatal(app.Listen(":" + loadConfig.HTTPPort))
	}

}
