package routes

import (
	"go_learning/internal/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine{
	router := gin.Default();
	router.GET("/healthCheck", handler.HealthCheckHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler));

	return router;
}