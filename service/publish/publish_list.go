package publish

import (
	"Douyin/common"
	"Douyin/repository"
	"Douyin/service/user"
)

// GetList 传入两个id，表示查询user的发布列表以及guest是否关注了user
func GetList(userId uint, guestId uint) ([]common.Video, error) {
	//1.查询作者信息，在GetUserInfo中已经实现了参数合法性检验
	author, err := user.GetUserInfo(userId, guestId)
	if err != nil {
		return nil, err
	}
	//2. 查询userId发布的所有视频
	videoList := repository.GetVideoListByUserId(userId)
	//3.组装成合适的格式并返回
	var publishList []common.Video
	for i := 0; i < len(videoList); i++ {
		video := common.Video{
			Id:            videoList[i].ID,
			Author:        author,
			PlayUrl:       videoList[i].PlayUrl,
			CoverUrl:      videoList[i].CoverUrl,
			FavoriteCount: repository.GetFavoritedCountByVideoId(videoList[i].ID),
			CommentCount:  repository.GetCommentCountById(videoList[i].ID),
			IsFavorite:    repository.IsBFavoriteA(videoList[i].ID, guestId),
			Title:         videoList[i].Title,
		}
		publishList = append(publishList, video)
	}
	return publishList, nil
}
