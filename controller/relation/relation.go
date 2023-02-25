package relation

import (
	"Douyin/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListResponse struct {
	common.Response
	List []common.User `json:"user_list"`
}

// Action no practical effect
func Action(c *gin.Context) {
	c.JSON(http.StatusOK, common.Response{StatusCode: 0})
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, ListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		List: []common.User{common.DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, ListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		List: []common.User{common.DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, ListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		List: []common.User{common.DemoUser},
	})
}
