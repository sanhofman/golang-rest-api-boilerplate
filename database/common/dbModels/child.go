package dbModels

import (
    "gorm.io/gorm"
)

type Child struct {
    gorm.Model
	Name            string             `json:"name"`
	Parent          string             `json:"parent"`
	CreatedBy       string             `json:"created_by"`
}
