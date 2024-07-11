package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoadRouter() *gin.Engine{
	router := gin.Default();
	UserRoute(router);


	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler));
	return router;
}