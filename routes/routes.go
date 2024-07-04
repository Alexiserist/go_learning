package routes

import "github.com/gin-gonic/gin"

func LoadRouter() *gin.Engine{
 	router := gin.Default();
	return router;
}