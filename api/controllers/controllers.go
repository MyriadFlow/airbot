package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	r.POST("/generate", Generate)
}

func Generate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
