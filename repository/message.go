package repository

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

// Message 存储消息，以gorm.Model.ID作为消息的id，表示user给to_user发送了消息，内容为content；
// CreateDate可通过gorm.Model.CreatedAt获取。
type Message struct {
	gorm.Model
	UserId   uint
	ToUserId uint
	Content  string
}

func SendMessage(message *Message) error {
	err := DB.Model(Message{}).Create(message).Error
	if err != nil {
		return errors.New("send message failed")
	}
	return nil
}

func GetMessageList(userId uint, toUserId uint, preTime time.Time) ([]Message, error) {
	var MessageList []Message
	result := DB.Model(Message{}).Where("((user_id = ? AND to_user_id = ?) OR (user_id = ? AND to_user_id = ?)) AND created_at > ?", userId, toUserId, toUserId, userId, preTime).Order("created_at asc").Find(&MessageList)
	if result.Error != nil {
		return nil, errors.New("get message list failed")
	}
	return MessageList, nil
}
