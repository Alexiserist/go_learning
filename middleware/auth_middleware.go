package middleware

import (
	"go_learning/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


type AuthMiddleware interface {
}

type authMiddleware struct{
	authService auth.AuthService
}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{
		authService: auth.NewAuthRepository(),
	}
}

func (m *authMiddleware) AuthorizationMiddleware(c *gin.Context){
	header := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(header,"Bearer ")

	if err := m.authService.ValidateToken(token); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"message":  "OK",
			"data": 	token,
		})
	}

}