package handler

import (
	"go_learning/internal/models"
	"go_learning/internal/services"
	"go_learning/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// @Summary Get all users
// @Description Get all users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} object{status=int,message=string,data=[]models.User}
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context){
	users,err := h.userService.GetAllUsers()
	if err != nil {
		utils.RespondWithStatusMessage(c,http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	utils.ResponseWithStatusNessageData(c,http.StatusOK,http.StatusText(http.StatusOK),users)
}

// @Summary Create users
// @Description Create users
// @Tags Users
// @Accept  json
// @Param body body models.User true "Request Body"
// @Produce  json
// @Success 200 {object} object{status=int,message=string,data=models.User}
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context){
	var users  models.User;
    if err := c.BindJSON(&users); err != nil {
		utils.RespondWithStatusMessage(c, http.StatusBadRequest, "Invalid request payload")
        return
    }
	log.Println(users);

	users,err := h.userService.CreateUser(users);
	if err != nil{
		utils.RespondWithStatusMessage(c,http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	utils.ResponseWithStatusNessageData(c,http.StatusOK,http.StatusText(http.StatusOK),users)
}