package feed

import (
	"Douyin/common"
	"Douyin/repository"
	"Douyin/service/user"
	"time"
)

func GetFeedList(guestId uint, lastTime int64) ([]common.Video, int64) {
	var newTime int64 = 0
	videoList := repository.GetVideoListByTime(time.Unix(lastTime/1000, 0), guestId)
	var retList []common.Video
	for _, x := range videoList {
		var tmp common.Video
		tmp.Id = x.ID
		tmp.PlayUrl = x.PlayUrl
		tmp.CoverUrl = x.CoverUrl
		tmp.Title = x.Title

		feedUser, _ := user.GetUserInfo(x.AuthorId, guestId)
		tmp.Author = feedUser

		tmp.FavoriteCount = repository.GetFavoritedCountByVideoId(tmp.Id)
		tmp.CommentCount = repository.GetCommentCountById(tmp.Id)
		tmp.IsFavorite = repository.IsBFavoriteA(tmp.Id, guestId)

		retList = append(retList, tmp)
		newTime = x.CreatedAt.Unix() * 1000
	}
	return retList, newTime
}
