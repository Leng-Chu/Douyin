package relation

import (
	"Douyin/common"
	"Douyin/service/relation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Action no practical effect
func Action(c *gin.Context) {

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

	var strToUserId = c.Query("to_user_id")
	toUserId, err := strconv.ParseUint(strToUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	err = relation.Action(userId, uint(toUserId), uint(actionType))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  "success!",
	})
}
