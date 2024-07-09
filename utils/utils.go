package utils

import "github.com/gin-gonic/gin"

type ApiStatusMessage struct {
	Status 	   int 		`json:"status"`
	Message    string 	`json:"message"`
	Data 	   interface{} `json:"data,omitempty"`
}

func RespondWithStatusMessage(c *gin.Context, code int, message string){
	body := ApiStatusMessage{
		Status: code,
		Message: message,
	}
	c.JSON(code,body)
	c.Abort();
}

func ResponseWithStatusNessageData(c *gin.Context, code int ,message string, data interface{}){
	body := ApiStatusMessage{
		Status: code,
		Message: message,
		Data: data,
	}
	c.JSON(code,body)
	c.Abort();
}
