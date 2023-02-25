package favorite

import (
	"Douyin/common"
	"Douyin/repository"
	"Douyin/service/user"
)

func GetList(userNewId uint, userHostId uint) []common.Video {

	var favoriteVideoList []common.Video
	favoriteVideoList = make([]common.Video, 0)

	videoList := GetVideoList(userNewId)
	for _, x := range videoList {
		var tmp common.Video
		tmp.Id = x.ID
		tmp.PlayUrl = x.PlayUrl
		tmp.CoverUrl = x.CoverUrl
		tmp.Title = x.Title

		favoriteUser, _ := user.GetUserInfo(x.AuthorId, userHostId)
		tmp.Author = favoriteUser

		tmp.FavoriteCount = repository.GetFavoritedCountByVideoId(tmp.Id)
		tmp.CommentCount = repository.GetCommentCountById(tmp.Id)
		tmp.IsFavorite = repository.IsBFavoriteA(tmp.Id, userHostId)

		favoriteVideoList = append(favoriteVideoList, tmp)
	}
	return favoriteVideoList

}

func GetVideoList(userId uint) []repository.Video {

	var favoriteList = repository.GetFavoriteByUserId(userId)

	videoList := make([]repository.Video, 0)

	for _, v := range favoriteList {
		var video = repository.GetVideoById(v.VideoId)
		videoList = append(videoList, video)
	}

	return videoList
}
