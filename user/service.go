package main

import (
	"errors"
	"mikro.host/models"
)

type UserService interface {
	Create(user models.CreateRequest) (models.User, error)
	GetAll() ([]models.User, error)
	Get(id string) (models.User, error)
}

type userService struct {
}

func (userService) Create(input models.CreateRequest) (user models.User, err error) {
	Db.Create(&models.User{
		Name: input.Name,
		Mail: input.Mail,
	}).Model(&models.User{}).First(&user, "mail = ?", input.Mail)

	return user, err
}

func (userService) GetAll() (users []models.User, err error) {
	Db.Model(&models.User{}).Find(&users)

	return users, nil
}

func (userService) Get(id string) (user models.User, err error) {
	Db.Model(&models.User{}).First(&user, "id = ?", id)
	if user.Mail == "" {
		return user, errors.New("not found")
	}

	return user, nil
}
