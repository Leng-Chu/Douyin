package comment

import (
	"Douyin/common"
	"Douyin/service/comment"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListResponse struct {
	common.Response
	CommentList []common.Comment `json:"comment_list,omitempty"`
}

func List(c *gin.Context) {
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	//调用service层评论函数
	commentList, err := comment.GetList(uint(videoId))
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
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  "查表成功！",
		},
		CommentList: commentList,
	})
	return
}
