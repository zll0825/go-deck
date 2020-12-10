package model

import (
	"go-deck/app/model/entity"
	"go-deck/app/util"
)

func (d *Dao) Login(user *entity.User) (*entity.User, error) {
	var data entity.User
	user.Password = util.HashPassword(user.Password)
	err := SystemDB().Where("username = ? AND password = ?", user.Username, user.Password).First(&data).Error
	return &data, err
}

func (d *Dao) FindUserByUserName(userName string) (*entity.User, error) {
	var user entity.User
	err := SystemDB().Where("username = ?", userName).First(&user).Error
	return &user, err
}

func (d *Dao) CreateUser(user *entity.User) error {
	user.Password = util.HashPassword(user.Password)
	err := SystemDB().Create(user).Error
	return err
}
