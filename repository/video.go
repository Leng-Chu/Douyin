package repository

import (
	"gorm.io/gorm"
	"time"
)

// Video 存储Video的基本信息，以gorm.Model.ID作为视频的id；
// 表中未存储点赞数、评论数和用户的具体信息，可通过调用其他文件中的接口获取；
// eg.在repository/comment文件中定义GetCommentCountById函数；调用service/user中的GetUserInfo函数得到用户信息。
type Video struct {
	gorm.Model
	AuthorId uint
	PlayUrl  string
	CoverUrl string
	Title    string
}

func InsertNewVideo(v *Video) error {
	if err := DB.Create(v).Error; err != nil {
		return err
	}
	return nil
}

func GetVideoListByUserId(id uint) []Video {
	var videoList []Video
	DB.Model(Video{}).Where("author_id=?", id).Find(&videoList)
	return videoList
}

func GetVideoListByTime(inputTime time.Time) []Video {
	var videoList []Video
	DB.Model(Video{}).Where("created_at < ?", inputTime).Order("created_at desc").Find(&videoList)
	return videoList
}

func GetVideoById(id uint) Video {
	var video Video
	DB.Model(Video{}).Where("id = ?", id).First(&video)
	return video
}

// GetWorkCountByUserId 根据用户id查询他发布了多少视频
func GetWorkCountByUserId(id uint) uint {
	var cnt int64
	DB.Model(Video{}).Where("author_id = ?", id).Count(&cnt)
	return uint(cnt)
}
