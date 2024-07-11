package routes

import (
	"go_learning/internal/handler"
	"go_learning/internal/repository"
	"go_learning/internal/services"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	userRepository := repository.NewUserRepository();
	userService := services.NewUserService(userRepository);
	userHandler := handler.NewUserHandler(userService);
	router.GET("/users", userHandler.GetUsers);
	router.GET("/users/:id", userHandler.GetUserByKey);
	router.POST("/users", userHandler.CreateUser);
	router.DELETE("/users/:id", userHandler.DeleteUser);
}