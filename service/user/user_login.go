package user

import (
	"Douyin/middleware/jwt"
	"Douyin/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func GetLoginInfo(username string, password string) (uint, string, error) {

	//1.合法性检验
	err := checkLogin(username, password)
	if err != nil {
		return 0, "", err
	}

	//2.获得id
	id := repository.GetIdByName(username)

	//3.获得token
	token, err := jwt.CreateToken(id)
	if err != nil {
		return 0, "", err
	}

	return id, token, nil
}

func checkLogin(name string, pass string) error {
	// 此处不需要验证用户名和密码过长过短的问题，直接验证用户是否存在以及密码是否正确即可

	// 得到密码的同时验证用户名是否存在
	hashedPass, err := repository.GetPassByName(name)
	if err != nil {
		return errors.New("user name doesn't exist")
	}

	// 验证密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass)); err != nil {
		return errors.New("incorrect password")
	}
	return nil
}
