package message

import (
	"Douyin/common"
	"Douyin/service/message"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

	strToUserId := c.Query("to_user_id")
	toUserId, err := strconv.ParseUint(strToUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	content := c.Query("content")

	err = message.Action(userId, uint(toUserId), content)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
	c.JSON(http.StatusOK, common.Response{StatusCode: 0})
}
