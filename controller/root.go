package controller

import (
	"SRHS-4cut-server/services"
	"github.com/gin-gonic/gin"
)

func RootController(r *gin.Engine) {
	r.POST("/", func(c *gin.Context) {
		services.GetImage(c)
	})

	r.GET("/", func(c *gin.Context) {

	})
}
