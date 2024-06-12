package SRHS_4cut_server

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
