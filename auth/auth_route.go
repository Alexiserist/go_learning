package auth

import (
	"github.com/gin-gonic/gin"
)

func authRoute(router *gin.Context){
	authRepository := NewAuthRepository();
	authService := NewAuthService(authRepository);
	authHandler := NewAuthHandler(authService);


	router.GET("/users", authHandler.GenerateToken);

}