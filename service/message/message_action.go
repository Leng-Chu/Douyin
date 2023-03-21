package message

import (
	"Douyin/repository"
	"errors"
	"fmt"
)

func Action(userId uint, toUserId uint, content string) error {
	// 检查用户id是否存在
	err := check(userId, toUserId)
	if err != nil {
		return err
	}
	newMessage := repository.Message{
		UserId:   userId,
		ToUserId: toUserId,
		Content:  content,
	}
	err = repository.SendMessage(&newMessage)
	if err != nil {
		return err
	}
	return nil
}

func check(userId uint, toUserId uint) error {
	if !repository.IsUserExistById(userId) {
		return errors.New(fmt.Sprintf("user(id=%v) doesn't exist", userId))
	}
	if !repository.IsUserExistById(toUserId) {
		return errors.New(fmt.Sprintf("user(id=%v) doesn't exist", toUserId))
	}
	return nil
}
