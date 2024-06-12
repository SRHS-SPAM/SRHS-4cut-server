package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get image",
	})
}
