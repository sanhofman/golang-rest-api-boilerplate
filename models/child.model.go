package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateChildRequest struct {
	Name      string    `json:"name" bson:"name" binding:"required"`
	Gender    string    `json:"gender" bson:"gender" binding:"required"`
	Parent    string    `json:"parent" bson:"parent" binding:"required"`
	Image     string    `json:"image,omitempty" bson:"image,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBChild struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Gender    string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Parent    string             `json:"parent,omitempty" bson:"parent,omitempty"`
	Image     string             `json:"image,omitempty" bson:"image,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateChild struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"title,omitempty" bson:"title,omitempty"`
	Gender    string             `json:"content,omitempty" bson:"content,omitempty"`
	Parent    string             `json:"parent,omitempty" bson:"parent,omitempty"`
	Image     string             `json:"image,omitempty" bson:"image,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

