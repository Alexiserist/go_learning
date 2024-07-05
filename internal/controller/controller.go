package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Health Check
// @Description Checks the health of the server
// @Tags health
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /healthCheck [get]
func HealthCheckController(g *gin.Context) {
	g.String(http.StatusOK, "OK")
}