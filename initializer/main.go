package initializer

import "github.com/maxahirwe/goleaf/models"

func init() {
	LoadEnv()
	LoadDB()
}

func Load() {
	DATABASE.AutoMigrate(&models.User{}) // performs migration for user table
}
