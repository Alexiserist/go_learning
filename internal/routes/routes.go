package routes

import (
	"go_learning/internal/handler"
	"go_learning/internal/repository"
	"go_learning/internal/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoadRouter() *gin.Engine{
	userRepository := repository.NewUserRepository();
	userService := services.NewUserService(userRepository);
	userHandler := handler.NewUserHandler(userService);


	router := gin.Default();
	router.GET("/healthCheck", handler.HealthCheckHandler);
	router.POST("/getTesting", handler.GetTesting);
	router.GET("/users", userHandler.GetUsers);
	router.POST("/users", userHandler.CreateUser);
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler));

	return router;
}