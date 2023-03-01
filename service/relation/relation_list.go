package relation

import (
	"Douyin/common"
	"Douyin/repository"
	"Douyin/service/user"
)

func GetFollowList(userId uint, guestId uint) ([]common.User, error) {
	//1. 查询userId关注的所有人
	userList := repository.GetFollowListById(userId)
	//2. 组装成合适的格式并返回
	var followList []common.User
	for i := 0; i < len(userList); i++ {
		u, err := user.GetUserInfo(userList[i].UserId, guestId)
		if err != nil {
			return nil, err
		}
		followList = append(followList, u)
	}
	return followList, nil
}

func GetFollowerList(userId uint, guestId uint) ([]common.User, error) {
	//1. 查询关注userId的所有人
	userList := repository.GetFollowerListById(userId)
	//2. 组装成合适的格式并返回
	var followerList []common.User
	for i := 0; i < len(userList); i++ {
		u, err := user.GetUserInfo(userList[i].FollowerId, guestId)
		if err != nil {
			return nil, err
		}
		followerList = append(followerList, u)
	}
	return followerList, nil
}

func GetFriendList(userId uint, guestId uint) ([]common.User, error) {
	//1. 查询userId的好友
	userList := repository.GetFriendListById(userId)
	//2. 组装成合适的格式并返回
	var friendList []common.User
	for i := 0; i < len(userList); i++ {
		u, err := user.GetUserInfo(userList[i].UserId, guestId)
		if err != nil {
			return nil, err
		}
		friendList = append(friendList, u)
	}
	return friendList, nil
}
