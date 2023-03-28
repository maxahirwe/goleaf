package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxahirwe/goleaf/controller"
	"github.com/maxahirwe/goleaf/initializer"
)

func init() {
	initializer.Load()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	apiV1 := r.Group("api/v1", gin.BasicAuth(gin.Accounts{"idt": "leaf"}))
	{
		apiV1.GET("/", controller.GetBaseUrl)
		apiV1.POST("/", controller.CreateUser)
		apiV1.GET("/:id", controller.GetUser)
		apiV1.POST("/all", controller.GetAllUsers)
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
