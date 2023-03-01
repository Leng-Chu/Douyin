package relation

import (
	"Douyin/repository"
	"errors"
)

func Action(userId uint, toUserId uint, actionType uint) (err error) {
	if userId == 0 {
		return errors.New("user not logged in")
	}
	if userId == toUserId {
		return errors.New("can't follow yourself")
	}
	isFollow := repository.IsBFollowA(toUserId, userId)
	if actionType == 1 && !isFollow {
		err := repository.AddFollow(toUserId, userId)
		if err != nil {
			return err
		}
	} else if actionType == 2 && isFollow {
		err := repository.DeleteFollow(toUserId, userId)
		if err != nil {
			return err
		}
	} else {
		return errors.New("invalid action")
	}
	return nil
}
