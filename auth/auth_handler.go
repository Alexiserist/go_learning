package auth

import (
	"go_learning/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// @Summary Get all users
// @Description Get all users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} object{status=int,message=string,data=[]models.User}
// @Router /users [get]
func (h *AuthHandler) GenerateToken(c *gin.Context) {
	token, err := h.authService.GenerateToken()
	if err != nil {
		utils.RespondWithStatusMessage(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	utils.ResponseWithStatusNessageData(c, http.StatusOK, http.StatusText(http.StatusOK), token)
}