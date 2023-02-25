package user

import (
	"Douyin/middleware/jwt"
	"Douyin/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

func GetRegisterInfo(username string, password string) (uint, string, error) {

	//1.合法性检验
	err := checkRegister(username, password)
	if err != nil {
		return 0, "", err
	}

	//2.新建用户
	id, err := createUser(username, password)
	if err != nil {
		return 0, "", err
	}

	//3.获得token
	token, err := jwt.CreateToken(id)
	if err != nil {
		return 0, "", err
	}

	return id, token, nil
}

func checkRegister(name string, pass string) error {
	if name == "" {
		return errors.New("user name is empty")
	}
	if len(name) > MaxUsernameLength {
		return errors.New("user name length exceeds the limit")
	}
	if len(pass) < MinPasswordLength {
		return errors.New("password is too short")
	}
	if len(pass) > MaxPasswordLength {
		return errors.New("password length exceeds the limit")
	}
	if repository.IsUserExistByName(name) {
		return errors.New("user name already exist")
	}
	return nil
}

func createUser(name string, password string) (uint, error) {
	// 对密码进行bcrypt加密操作
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	newPass := string(hashedPass)

	newUser := repository.User{
		Name:     name,
		Password: newPass,
	}
	err = repository.InsertNewUser(&newUser)
	if err != nil {
		return 0, err
	}
	return newUser.ID, nil
}
