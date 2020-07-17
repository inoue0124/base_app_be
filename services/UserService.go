package services

import (
	"errors"

	"github.com/inoue0124/salon-be/db"
	"github.com/inoue0124/salon-be/forms"
	"github.com/inoue0124/salon-be/models"
)

type UserService struct{}

type User models.User

func (s UserService) Signup(userPayload forms.UserSignup) (User, error) {
	db := db.GetDB()
	db.AutoMigrate(&User{})
	var user User
	db.First(&user, "email=?", userPayload.Email)
	if user.Email != "" {
		return user, errors.New("The user is already exist")
	}
	user = User{
		Email: userPayload.Email,
	}
	db.Create(&user)
	return user, nil
}

func (s UserService) GetByID(id string) (User, error) {
	db := db.GetDB()
	var user User
	db.First(&user, id)
	if user.ID == 0 {
		return user, errors.New("The user is not exist")
	}
	return user, nil
}
