package handler

import (
	"net/http"

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