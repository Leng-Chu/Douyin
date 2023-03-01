package repository

import (
	"errors"
	"gorm.io/gorm"
)

// Relation 存储关注信息，表示follower关注了user。
type Relation struct {
	gorm.Model
	UserId     uint
	FollowerId uint
}

func GetFollowCountById(id uint) uint {
	var cnt int64
	DB.Model(Relation{}).Where("follower_id = ?", id).Count(&cnt)
	return uint(cnt)
}

func GetFollowerCountById(id uint) uint {
	var cnt int64
	DB.Model(Relation{}).Where("user_id = ?", id).Count(&cnt)
	return uint(cnt)
}

// IsBFollowA 用户B是否关注用户A
func IsBFollowA(idA uint, idB uint) bool {
	if idA == idB || idB == 0 {
		return false
	}
	var isFollow Relation
	result := DB.Model(Relation{}).Where("follower_id = ? AND user_id = ?", idB, idA).First(&isFollow)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func AddFollow(idA uint, idB uint) error {
	newFollow := Relation{
		FollowerId: idB,
		UserId:     idA,
	}
	if err := DB.Create(&newFollow).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFollow(idA uint, idB uint) error {
	oldFollow := Relation{}
	if err := DB.Where("follower_id = ? AND user_id = ?", idB, idA).Take(&oldFollow).Error; err != nil {
		return err
	}
	if err := DB.Delete(&oldFollow).Error; err != nil {
		return err
	}
	return nil
}

func GetFollowListById(id uint) []Relation {
	var followList []Relation
	DB.Model(Relation{}).Where("follower_id=?", id).Find(&followList)
	return followList
}

func GetFollowerListById(id uint) []Relation {
	var followerList []Relation
	DB.Model(Relation{}).Where("user_id=?", id).Find(&followerList)
	return followerList
}

func GetFriendListById(id uint) []Relation {
	followList := GetFollowListById(id)
	followerList := GetFollowerListById(id)
	var friendList []Relation
	for _, following := range followList {
		for _, follower := range followerList {
			if following.UserId == follower.FollowerId {
				friendList = append(friendList, following)
				break
			}
		}
	}
	return friendList
}
