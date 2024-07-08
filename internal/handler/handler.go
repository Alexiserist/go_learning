package handler

import (
	"net/http"
	"go_learning/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary Health Check
// @Description Checks the health of the server
// @Tags Health Check
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /healthCheck [get]
func HealthCheckHandler(g *gin.Context) {
	g.String(http.StatusOK, "OK")
}


// @Summary Testing
// @Description Testing method get post delete update
// @Tags Testing
// @Accept json
// @Produce json
// @Param body body models.Testing true "Request Body"
// @Success 200 {object} models.Testing
// @Router /getTesting [post]
func GetTesting(g *gin.Context){
	var model  models.Testing;
	if err := g.ShouldBindBodyWithJSON(model); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
		return;
	}
	g.String(http.StatusOK, "OK");

}