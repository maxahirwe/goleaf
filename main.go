package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxahirwe/goleaf/controller"
	"github.com/maxahirwe/goleaf/initializer"
)

func init() {
	initializer.Load()
}

func main() {
	r := gin.Default()
	r.GET("/", controller.GetBaseUrl)
	r.POST("/", controller.CreateUser)
	r.GET("/:id", controller.GetUser)
	r.POST("/all", controller.GetAllUsers)
	r.Run()
}
