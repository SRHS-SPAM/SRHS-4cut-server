package services

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	"net/http"
)

var currentImage *imageupload.Image

func PostImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = c.SaveUploadedFile(file, "/imageData"+file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
