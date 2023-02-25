package comment

import (
	"Douyin/common"
	"Douyin/service/comment"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ActionResponse struct {
	common.Response
	Comment common.Comment `json:"comment,omitempty"`
}

func Action(c *gin.Context) {

	//从token中解析出id，得到的是当前登录用户的id
	idToken, _ := c.Get("user_id")
	userId, ok := idToken.(uint)
	if !ok {
		c.JSON(http.StatusOK, ActionResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "invalid user id",
			},
		})
		return
	}

	var strVideoId = c.Query("video_id")
	videoId, err := strconv.ParseUint(strVideoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ActionResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	var strActionType = c.Query("action_type")
	actionType, err := strconv.ParseUint(strActionType, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ActionResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	if actionType == 1 {
		content := c.Query("comment_text")
		newComment, err := comment.Insert(userId, uint(videoId), content)
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		c.JSON(http.StatusOK, ActionResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "add comment successfully",
			},
			Comment: newComment,
		})
	} else {
		commentId, err := strconv.ParseUint(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		err = comment.Delete(uint(commentId))
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				Response: common.Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
		c.JSON(http.StatusOK, ActionResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "delete successfully",
			},
		})
	}
}
