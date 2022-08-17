package services

import (
	"github.com/wpcodevo/golang-mongodb/database/common/dbModels"

	"github.com/wpcodevo/golang-mongodb/models"
)

type AuthService interface {
	SignUpUser(*models.SignUpInput) (dbModels.User, error)
	SignInUser(*models.SignInInput) (dbModels.User, error)
}
