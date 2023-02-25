package repository

import (
	"errors"
	"gorm.io/gorm"
)

// Favorite 存储点赞信息，表示user点赞了video。
type Favorite struct {
	gorm.Model
	UserId  uint
	VideoId uint
}

// GetFavoritedCountByVideoId 通过视频id查询被赞的数量
func GetFavoritedCountByVideoId(id uint) uint {
	var cnt int64
	DB.Model(Favorite{}).Where("video_id = ?", id).Count(&cnt)
	return uint(cnt)
}

// GetFavoriteByUserId 通过用户id查询点赞的所有视频
func GetFavoriteByUserId(id uint) []Favorite {
	var favoriteList []Favorite
	DB.Model(Favorite{}).Where("user_id = ?", id).Find(&favoriteList)
	return favoriteList
}

// GetFavoriteCountByUserId 通过用户id查询点赞的数量
func GetFavoriteCountByUserId(id uint) uint {
	var cnt int64
	DB.Model(Favorite{}).Where("user_id = ?", id).Count(&cnt)
	return uint(cnt)
}

// GetTotalFavoritedByUserId 通过用户id查询被点赞的数量
func GetTotalFavoritedByUserId(id uint) uint {
	var cnt int64
	DB.Model(Favorite{}).Joins("JOIN videos ON favorites.video_id = videos.id").Where("videos.author_id = ?", id).Count(&cnt)
	return uint(cnt)
}

// IsBFavoriteA 用户B是否给视频A点赞
func IsBFavoriteA(idA uint, idB uint) bool {
	if idB == 0 {
		return false
	}
	var isFavorite Favorite
	result := DB.Model(Favorite{}).Where("user_id = ? AND video_id = ?", idB, idA).First(&isFavorite)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func AddFavorite(idA uint, idB uint) error {
	newFavorite := Favorite{
		UserId:  idB,
		VideoId: idA,
	}
	if err := DB.Create(&newFavorite).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFavorite(idA uint, idB uint) error {
	oldFavorite := Favorite{}
	if err := DB.Where("user_id = ? AND video_id = ?", idB, idA).Take(&oldFavorite).Error; err != nil {
		return err
	}
	if err := DB.Delete(&oldFavorite).Error; err != nil {
		return err
	}
	return nil
}
