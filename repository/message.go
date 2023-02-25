package repository

import "gorm.io/gorm"

// Message 存储消息，以gorm.Model.ID作为消息的id，表示user给to_user发送了消息，内容为content；
// CreateDate可通过gorm.Model.CreatedAt获取。
type Message struct {
	gorm.Model
	UserId   uint
	ToUserId uint
	Content  string
}
