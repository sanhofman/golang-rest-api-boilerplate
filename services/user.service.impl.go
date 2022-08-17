package services

import (
	"context"
	"errors"

    "gorm.io/gorm"

	"github.com/wpcodevo/golang-mongodb/database/common/dbModels"
)

type UserServiceImpl struct {
	db             *gorm.DB
	ctx            context.Context
}

func NewUserServiceImpl(db *gorm.DB, ctx context.Context) UserService {
	return &UserServiceImpl{db, ctx}
}

func (us *UserServiceImpl) FindUserById(id string) (dbModels.User, error) {
    var user dbModels.User

    us.db.Where("id = ?", id).First(&user)
    if user.Email == "" {
        return user, errors.New("user not found")
    }

	return user, nil
}

func (us *UserServiceImpl) FindUserByEmail(email string) (dbModels.User, error) {
    var user dbModels.User

    us.db.Where("email = ?", email).First(&user)
    if user.Email == "" {
        return user, errors.New("user not found")
    }

	return user, nil
}

func (uc *UserServiceImpl) UpdateUserById(id string, field string, value string) (dbModels.User, error) {
    var user dbModels.User

    uc.db.Where("uuid = ?", id).First(&user)
    if user.Email != "" {
        return user, errors.New("user not found")
    }

    uc.db.Model(&user).Update(field, value)

	return user, nil
}

func (uc *UserServiceImpl) UpdateOne(field string, value interface{}) (dbModels.User, error) {
// 	query := bson.D{{Key: field, Value: value}}
// 	update := bson.D{{Key: "$set", Value: bson.D{{Key: field, Value: value}}}}
// 	result, err := uc.collection.UpdateOne(uc.ctx, query, update)
//
// 	fmt.Print(result.ModifiedCount)
// 	if err != nil {
// 		fmt.Print(err)
// 		return &models.DBResponse{}, err
// 	}
//
// 	return &models.DBResponse{}, nil

    // @TODO:: what is this?
    var user dbModels.User
    return user, nil
}
