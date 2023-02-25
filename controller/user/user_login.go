package user

import (
	"Douyin/common"
	"Douyin/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginResponse struct {
	common.Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Login(c *gin.Context) {

	// 参数提取部分
	username := c.Query("username")
	password := c.Query("password")

	// 调用service层得到响应
	id, token, err := user.GetLoginInfo(username, password)

	// 将调用service层得到的Info与Response打包，返回响应
	if err != nil {
		c.JSON(http.StatusOK, LoginResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, LoginResponse{
		Response: common.Response{StatusCode: 0},
		UserId:   id,
		Token:    token,
	})
}
