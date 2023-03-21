package favorite

import (
	"Douyin/repository"
	"errors"
)

func Action(userId uint, videoId uint, actionType uint) (err error) {
	if userId == 0 {
		return errors.New("user not logged in")
	}
	isFavorite := repository.IsBFavoriteA(videoId, userId)
	if actionType == 1 && !isFavorite {
		// 点赞操作
		err := repository.AddFavorite(videoId, userId)
		if err != nil {
			return err
		}
	} else if actionType == 2 && isFavorite {
		// 取消点赞操作
		err := repository.DeleteFavorite(videoId, userId)
		if err != nil {
			return err
		}
	} else {
		return errors.New("invalid action")
	}
	return nil
}
