package repository

import (
	"errors"

	"gorm.io/gorm"
)

// Comment 存储评论信息，以gorm.Model.ID作为评论的id，表示user评论了video，内容为content；
// CreateDate可通过gorm.Model.CreatedAt获取。
type Comment struct {
	gorm.Model
	UserId  uint
	VideoId uint
	Content string
}

// GetCommentCountById 通过视频id查询评论的数量
func GetCommentCountById(id uint) uint {
	var cnt int64
	//数据库中查询评论数量
	DB.Model(Comment{}).Where("video_id = ?", id).Count(&cnt)
	return uint(cnt)
}

// InsertComment 发表评论
func InsertComment(comment *Comment) error {
	err := DB.Model(Comment{}).Create(comment).Error
	if err != nil {
		return errors.New("create comment failed")
	}
	return nil
}

// DeleteComment 删除评论
func DeleteComment(id uint) error {
	oldComment := Comment{}
	if err := DB.Where("id = ? ", id).Take(&oldComment).Error; err != nil {
		return err
	}
	if err := DB.Delete(&oldComment).Error; err != nil {
		return err
	}
	return nil
}

// GetCommentList 获取评论列表
func GetCommentList(id uint) ([]Comment, error) {
	var CommentList []Comment
	result := DB.Model(Comment{}).Where("video_id = ?", id).Find(&CommentList)
	if result.Error != nil {
		return nil, errors.New("get comment list failed")
	}
	return CommentList, nil
}
