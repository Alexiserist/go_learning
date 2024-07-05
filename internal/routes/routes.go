package routes

import (
	"go_learning/internal/controller"

	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine{
	router := gin.Default();
	router.GET("/healthCheck", controller.HealthCheckController)

	return router;
}