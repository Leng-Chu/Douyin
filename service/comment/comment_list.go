package comment

import (
	"Douyin/common"
	"Douyin/repository"
	"Douyin/service/user"
)

func GetList(videoId uint) ([]common.Comment, error) {
	commentList, err := repository.GetCommentList(videoId)
	if err != nil {
		return nil, err
	}
	var retList []common.Comment
	for _, x := range commentList {
		u, err := user.GetUserInfo(x.UserId, x.UserId)
		if err != nil {
			return nil, err
		}
		comment := common.Comment{
			Id:         x.ID,
			User:       u,
			Content:    x.Content,
			CreateDate: x.CreatedAt.Format("01-02"),
		}
		retList = append(retList, comment)
	}
	return retList, nil
}
