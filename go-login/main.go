package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "go login!",
		})
	})

	r.Run(":8081")
}
