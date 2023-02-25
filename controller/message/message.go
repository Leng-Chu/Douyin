package message

import (
	"Douyin/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var tempChat = map[string][]common.Message{}

var messageIdSequence = uint64(1)

type ChatResponse struct {
	common.Response
	MessageList []common.Message `json:"message_list"`
}

// Action no practical effect
func Action(c *gin.Context) {
	user := common.DemoUser
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	userIdB, _ := strconv.Atoi(toUserId)
	chatKey := genChatKey(user.Id, uint(userIdB))

	atomic.AddUint64(&messageIdSequence, 1)
	curMessage := common.Message{
		Id:         uint(messageIdSequence),
		Content:    content,
		CreateTime: time.Now().Format(time.Kitchen),
	}

	if messages, exist := tempChat[chatKey]; exist {
		tempChat[chatKey] = append(messages, curMessage)
	} else {
		tempChat[chatKey] = []common.Message{curMessage}
	}
	c.JSON(http.StatusOK, common.Response{StatusCode: 0})
}

// Chat all users have same follow list
func Chat(c *gin.Context) {
	user := common.DemoUser
	toUserId := c.Query("to_user_id")

	userIdB, _ := strconv.Atoi(toUserId)
	chatKey := genChatKey(user.Id, uint(userIdB))

	c.JSON(http.StatusOK, ChatResponse{Response: common.Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
}

func genChatKey(userIdA uint, userIdB uint) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
