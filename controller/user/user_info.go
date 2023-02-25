package user

import (
	"Douyin/common"
	"Douyin/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type InfoResponse struct {
	common.Response
	User common.User `json:"user"`
}

func Info(c *gin.Context) {

	//从token中解析出id，得到的是当前登录用户的id
	idToken, _ := c.Get("user_id")
	guestId, ok := idToken.(uint)
	if !ok {
		c.JSON(http.StatusOK, InfoResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "invalid user id",
			},
		})
		return
	}

	//调用query方法得到要查询的用户id
	idStr := c.Query("user_id")
	userId, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, InfoResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 因为要返回is_follow的信息，所以接口传入了两个id
	userInfo, err := user.GetUserInfo(uint(userId), guestId)
	if err != nil {
		c.JSON(http.StatusOK, InfoResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, InfoResponse{
		Response: common.Response{StatusCode: 0},
		User:     userInfo,
	})
	return
}
