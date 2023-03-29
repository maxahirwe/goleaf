package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maxahirwe/goleaf/initializer"
	"github.com/maxahirwe/goleaf/models"
)

// base url endpoint
func GetBaseUrl(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "users endpoints",
	})
}

// create user endpoint
func CreateUser(c *gin.Context) {
	minYear := 1850
	var UserBody struct {
		Name       string `json:"Name"  binding:"required,min=2"`
		SignupTime int64  `json:"SignupTime"  binding:"required,numeric"`
	}
	if err := c.ShouldBindJSON(&UserBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}
	signupTime := time.UnixMilli(UserBody.SignupTime)
	if signupTime.Year() < minYear {
		c.JSON(http.StatusBadRequest, gin.H{"validation error": "Key: 'SignupTime' Error:Field validation for 'SignupTime' failed on min time year of 1850"})
		return
	}
	user := models.User{Name: UserBody.Name, SignupTime: signupTime.UnixMilli()}
	initializer.DATABASE.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"status": "user created", "data": user})
}

// get user endpoint
func GetUser(c *gin.Context) {
	id := c.Param("id") // get id param
	var user models.User
	res := initializer.DATABASE.First(&user, id) // find a user by param id
	if res.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"user": nil,
		})
	}

}

// get all users endpoint
func GetAllUsers(c *gin.Context) {
	var users []models.User
	// can add pagination
	initializer.DATABASE.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
