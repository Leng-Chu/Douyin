package favorite

import (
	"Douyin/common"
	"Douyin/service/favorite"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Action 进行点赞操作
func Action(c *gin.Context) {

	//从token中解析出id，得到的是当前登录用户的id
	idToken, _ := c.Get("user_id")
	userId, ok := idToken.(uint)
	if !ok {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "invalid user id",
		})
		return
	}

	var strActionType = c.Query("action_type")
	actionType, err := strconv.ParseUint(strActionType, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	var strVideoId = c.Query("video_id")
	VideoId, err := strconv.ParseUint(strVideoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	err = favorite.Action(userId, uint(VideoId), uint(actionType))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 0,
			StatusMsg:  "success!",
		})
	}

}
