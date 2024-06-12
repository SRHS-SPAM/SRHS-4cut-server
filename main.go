package main

import (
	"SRHS-4cut-server/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	controller.RootController(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
