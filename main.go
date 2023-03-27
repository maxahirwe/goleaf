package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxahirwe/goleaf/initializer"
)

func init() {
	initializer.Load()
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user challenge",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}