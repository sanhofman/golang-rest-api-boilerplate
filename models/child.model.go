package models

import (
	"time"
)

type CreateChildRequest struct {
	Name      string    `json:"name" binding:"required"`
	Gender    string    `json:"gender" binding:"required"`
	Parent    string    `json:"parent" binding:"required"`
	Image     string    `json:"image,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateChild struct {
	Name      string             `json:"name,omitempty" binding:"required"`
	Gender    string             `json:"gender,omitempty"`
	Parent    string             `json:"parent,omitempty"`
	Image     string             `json:"image,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}
