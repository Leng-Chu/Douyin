package feed

import (
	"Douyin/common"
	"Douyin/service/feed"
	"github.com/gin-gonic/gin"
	_ "github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

type ListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// List same demo video list for every request
// 注意：如果用户已登录，推送的视频流中不应该包含本人发布的视频
func List(c *gin.Context) {
	//从token中解析出id，得到的是当前登录用户的id
	idToken, _ := c.Get("user_id")
	guestId, ok := idToken.(uint)
	if !ok {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "invalid user id",
			},
		})
		return
	}

	var strLastTime = c.Query("latest_time")
	lastTime, err := strconv.ParseInt(strLastTime, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	feedList, newTime := feed.GetFeedList(guestId, lastTime)
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	if len(feedList) > 0 {
		c.JSON(http.StatusOK, ListResponse{
			Response:  common.Response{StatusCode: 0}, //成功
			VideoList: feedList,
			NextTime:  newTime,
		})
	} else {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{StatusCode: 0}, //成功
			NextTime: 0,                              //重新循环
		})
	}
}
