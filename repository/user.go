package repository

import (
	"errors"
	"gorm.io/gorm"
)

// User 存储User的基本信息，以gorm.Model.ID作为用户的id，name为明文，password加密后存储在数据库。
type User struct {
	gorm.Model
	Name     string
	Password string
}

func IsUserExistByName(name string) bool {
	var userExist = &User{}
	result := DB.Model(User{}).Where("name=?", name).First(&userExist)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 如果返回记录不存在的错误，说明没有名为name的用户
		return false
	}
	return true
}

func IsUserExistById(id uint) bool {
	var userExist = &User{}
	result := DB.Model(User{}).Where("id=?", id).First(&userExist)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func InsertNewUser(u *User) error {
	if err := DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func GetIdByName(name string) uint {
	var user = &User{}
	DB.Model(User{}).Where("name=?", name).First(&user)
	return user.ID
}

func GetNameById(id uint) string {
	var user = &User{}
	DB.Model(User{}).Where("id=?", id).First(&user)
	return user.Name
}

func GetPassByName(name string) (string, error) {
	var userId = &User{}
	err := DB.Model(User{}).Where("name=?", name).First(&userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}
	return userId.Password, nil
}
