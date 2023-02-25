package repository

import "gorm.io/gorm"

// Relation 存储关注信息，表示follower关注了user。
type Relation struct {
	gorm.Model
	UserId     uint
	FollowerId uint
}

func GetFollowCountById(id uint) uint {
	return 0
}

func GetFollowerCountById(id uint) uint {
	return 0
}

// IsBFollowA 用户B是否关注用户A
func IsBFollowA(idA uint, idB uint) bool {
	if idA == idB || idB == 0 {
		return false
	}
	// 未完善
	return false
}
