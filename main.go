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
	apiv1 := r.Group("api/v1", gin.BasicAuth(gin.Accounts{"idt": "leaf"}))
	{
		apiv1.GET("/", controller.GetBaseUrl)
		apiv1.POST("/", controller.CreateUser)
		apiv1.GET("/:id", controller.GetUser)
		apiv1.POST("/all", controller.GetAllUsers)
	}
	r.Run()
}
