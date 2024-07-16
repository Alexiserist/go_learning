package routes

import (
	"go_learning/database"
	"go_learning/internal/handler"
	"go_learning/internal/repository"
	"go_learning/internal/services"
	"go_learning/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	authMiddleware := middleware.NewAuthMiddleware();

	userRepository := repository.NewUserRepository(database.DB);
	userService := services.NewUserService(userRepository);
	userHandler := handler.NewUserHandler(userService);

	users := router.Group("/users");
	users.Use(authMiddleware.AuthorizationMiddleware);
	users.GET("/", userHandler.GetUsers);
	users.GET("/:id", userHandler.GetUserByKey);
	users.POST("/", userHandler.CreateUser);
	users.DELETE("/:id", userHandler.DeleteUser);
}