package services

import (
	"context"
	"errors"
	"strings"
	"time"

    "gorm.io/gorm"

	"github.com/wpcodevo/golang-mongodb/database/common/dbModels"

	"github.com/wpcodevo/golang-mongodb/models"
	"github.com/wpcodevo/golang-mongodb/utils"
)

type AuthServiceImpl struct {
	db             *gorm.DB
	ctx            context.Context
}

func NewAuthService(db *gorm.DB, ctx context.Context) AuthService {
	return &AuthServiceImpl{db, ctx}
}

func (uc *AuthServiceImpl) SignUpUser(User *models.SignUpInput) (dbModels.User, error) {
    var user dbModels.User

    // Check existence.
    uc.db.Where("email = ?", User.Email).First(&user)
    if user.Email != "" {
        return user, errors.New("user already exists")
    }

	user.CreatedAt = time.Now()
	user.UpdatedAt = User.CreatedAt
	user.Email = strings.ToLower(User.Email)
//	user.PasswordConfirm = ""
	user.Verified = false
//	user.Role = "user"

	hashedPassword, _ := utils.HashPassword(User.Password)
	user.Password = hashedPassword

    if result := uc.db.Create(&user); result.Error != nil {
        return user, errors.New("User creation failed")
    }

	return user, nil
}

func (uc *AuthServiceImpl) SignInUser(*models.SignInInput) (dbModels.User, error) {
    var user dbModels.User

	return user, nil
}
