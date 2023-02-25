package favorite

import (
	"Douyin/repository"
	"errors"
)

func Action(userId uint, VideoId uint, actionType uint) (err error) {
	if userId == 0 {
		return errors.New("user not logged in")
	}
	isFavorite := repository.IsBFavoriteA(VideoId, userId)
	if actionType == 1 && !isFavorite {
		// 点赞操作
		err := repository.AddFavorite(VideoId, userId)
		if err != nil {
			return err
		}
	} else if actionType == 0 && isFavorite {
		// 取消点赞操作
		err := repository.DeleteFavorite(VideoId, userId)
		if err != nil {
			return err
		}
	} else {
		return errors.New("invalid action")
	}
	return nil
}
