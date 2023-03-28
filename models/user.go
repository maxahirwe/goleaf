package models

import "gorm.io/gorm"

// User Struct/Model
type User struct {
	gorm.Model        // id, createdAt, updatedAt deletedAt will be auto populated
	Name       string `gorm:"not null"`
	SignupTime int64  `gorm:"autoUpdateTime:milli;not null"` // Use unix milli seconds as updating time
}
