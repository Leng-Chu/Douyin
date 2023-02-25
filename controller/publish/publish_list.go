package publish

import (
	"Douyin/common"
	"Douyin/service/publish"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListResponse struct {
	common.Response
	List []common.Video `json:"video_list"`
}

// List 查询user发布的视频列表
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

	//调用query方法得到要查询的用户id
	idStr := c.Query("user_id")
	userId, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 因为要返回is_follow和is_favorite的信息，所以接口传入了两个id
	publishList, err := publish.GetList(uint(userId), guestId)
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, ListResponse{
		Response: common.Response{StatusCode: 0},
		List:     publishList,
	})
}
