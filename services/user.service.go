package services

import (
	"github.com/wpcodevo/golang-mongodb/database/common/dbModels"
)

type UserService interface {
	FindUserById(id string) (dbModels.User, error)
	FindUserByEmail(email string) (dbModels.User, error)
	UpdateUserById(id string, field string, value string) (dbModels.User, error)
	UpdateOne(field string, value interface{}) (dbModels.User, error)
}
