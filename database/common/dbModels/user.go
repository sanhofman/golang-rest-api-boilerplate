package dbModels

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
	Uuid            string             `json:"uuid"`
	Name            string             `json:"name"`
	Email           string             `json:"email"`
	Password        string             `json:"password"`
	Verified        bool               `json:"verified"`
}
