package message

import (
	"Douyin/common"
	"Douyin/service/message"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ChatResponse struct {
	common.Response
	MessageList []common.Message `json:"message_list"`
}

func Chat(c *gin.Context) {
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
	strPreTime := c.Query("pre_msg_time")
	preTime, err := strconv.ParseInt(strPreTime, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	messageList, err := message.GetMessageList(userId, uint(toUserId), preTime)
	if err != nil {
		c.JSON(http.StatusOK, ChatResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, ChatResponse{
		Response:    common.Response{StatusCode: 0}, //成功
		MessageList: messageList,
	})
}
