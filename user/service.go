package main

import (
	"mikro.host/common"
	"mikro.host/models"
)

type UserService interface {
	Create(user models.CreateRequest) (models.User, error)
	GetAll() ([]models.User, error)
	Get(id uint) (models.User, error)
	Update(changes models.UpdateRequest) (models.User, error)
	Delete(id uint) error
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

func (userService) Get(id uint) (user models.User, err error) {
	Db.Model(&models.User{}).First(&user, "id = ?", id)
	if user.Mail == "" {
		return user, common.NotFound
	}

	return user, nil
}

func (userService) Update(changes models.UpdateRequest) (user models.User, err error) {
	Db.Model(&models.User{}).Updates(models.User{}).First(&user, "")
	if user.ID == 0 {
		return user, common.NotFound
	}

	return user, nil
}

func (u userService) Delete(id uint) error {
	if err := Db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
