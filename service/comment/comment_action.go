package comment

import (
	"Douyin/common"
	"Douyin/repository"
	"Douyin/service/user"
)

func Insert(userId uint, videoId uint, content string) (common.Comment, error) {
	newComment := repository.Comment{
		UserId:  userId,
		VideoId: videoId,
		Content: content,
	}
	err := repository.InsertComment(&newComment)
	if err != nil {
		return common.Comment{}, err
	}
	u, _ := user.GetUserInfo(userId, userId)
	retComment := common.Comment{
		Id:         newComment.ID,
		User:       u,
		Content:    content,
		CreateDate: newComment.CreatedAt.Format("01-02"),
	}
	return retComment, nil
}

func Delete(id uint) error {
	err := repository.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}
