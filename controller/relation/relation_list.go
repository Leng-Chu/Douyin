package relation

import (
	"Douyin/common"
	"Douyin/service/relation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListResponse struct {
	common.Response
	List []common.User `json:"user_list"`
}

// FollowList 返回user的关注列表
func FollowList(c *gin.Context) {
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

	var strUserId = c.Query("user_id")
	userId, err := strconv.ParseUint(strUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	followList, err := relation.GetFollowList(uint(userId), guestId)
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
		List:     followList,
	})
}

// FollowerList 返回user的粉丝列表
func FollowerList(c *gin.Context) {
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

	var strUserId = c.Query("user_id")
	userId, err := strconv.ParseUint(strUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	followerList, err := relation.GetFollowerList(uint(userId), guestId)
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
		List:     followerList,
	})
}

// FriendList 返回user的互关列表
func FriendList(c *gin.Context) {
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
	var strUserId = c.Query("user_id")
	userId, err := strconv.ParseUint(strUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	friendList, err := relation.GetFriendList(uint(userId), guestId)
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
		List:     friendList,
	})
}
