package services

import "github.com/wpcodevo/golang-mongodb/models"

type ChildService interface {
	CreateChild(*models.CreateChildRequest) (*models.DBChild, error)
	UpdateChild(string, *models.UpdateChild) (*models.DBChild, error)
	FindChildById(string) (*models.DBChild, error)
	FindChildren(page int, limit int) ([]*models.DBChild, error)
	DeleteChild(string) error
}
