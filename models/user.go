package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `gorm:"not null"`
	SignupTime int64  `gorm:"autoUpdateTime:milli;not null"` // Use unix milli seconds as updating time
}
