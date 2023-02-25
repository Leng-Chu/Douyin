package favorite

import (
	"Douyin/common"
	"Douyin/service/favorite"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListResponse struct {
	common.Response
	List []common.Video `json:"video_list"`
}

// List all users have same favorite video list
func List(c *gin.Context) {

	//从token中解析出id，得到的是当前登录用户的id
	idToken, _ := c.Get("user_id")
	userHostId, ok := idToken.(uint)
	if !ok {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "invalid user id",
			},
		})
		return
	}

	var strUserId = c.Query("user_id")
	userId, err := strconv.ParseUint(strUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	userNewId := uint(userId)
	if userNewId == 0 {
		userNewId = userHostId
	}

	favoriteList := favorite.GetList(userNewId, userHostId)

	if err != nil {
		c.JSON(http.StatusBadRequest, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "查表失败！",
			},
			List: nil,
		})
	} else {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "查表成功！",
			},
			List: favoriteList,
		})
	}
}
