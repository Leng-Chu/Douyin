package user

import (
	"Douyin/common"
	"Douyin/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterResponse struct {
	common.Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {

	// 参数提取部分
	username := c.Query("username")
	password := c.Query("password")

	// 调用service层得到响应
	id, token, err := user.GetRegisterInfo(username, password)

	// 将调用service层得到的Info与Response打包，返回响应
	if err != nil {
		c.JSON(http.StatusOK, RegisterResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, RegisterResponse{
		Response: common.Response{StatusCode: 0},
		UserId:   id,
		Token:    token,
	})
	return
}
