package services

import (
	"github.com/wpcodevo/golang-mongodb/database/common/dbModels"

	"github.com/gin-gonic/gin"

	"github.com/wpcodevo/golang-mongodb/models"
)

type ChildService interface {
	CreateChild(*models.CreateChildRequest, *gin.Context) (dbModels.Child, error)
	UpdateChild(string, *models.UpdateChild) (dbModels.Child, error)
	FindChildById(string) (dbModels.Child, error)
	FindChildren(page int, limit int) ([]dbModels.Child, error)
	DeleteChild(string) error
}
