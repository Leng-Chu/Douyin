package publish

import (
	"Douyin/common"
	"Douyin/service/publish"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
	"path/filepath"
)

// Action save upload file to data directory
func Action(c *gin.Context) {

	// 1. 获取视频
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 2. 获取用户
	idToken, _ := c.Get("user_id")
	id, ok := idToken.(uint)
	if !ok {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "invalid user id",
		})
		return
	}

	// 3. 生成视频名并将视频保存到本地
	name := uuid.NewV4().String()             // 生成随机的uuid作为视频名
	playSuffix := filepath.Ext(data.Filename) // 得到视频文件类型
	savePath := "./data/"
	// 如果文件夹不存在则创建
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		if err = os.MkdirAll(savePath, os.ModePerm); err != nil {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
	}
	// 将视频存储在本地
	if err := c.SaveUploadedFile(data, filepath.Join(savePath, name+playSuffix)); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 4. 获取一帧图片作为封面并保存到本地
	coverSuffix, err := publish.GetVideoCover(savePath, name, playSuffix)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 5. 传入用户id、标题、视频名、图片名的信息，更新数据库
	err = publish.UploadVideo(id, title, name+playSuffix, name+coverSuffix)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 6. 上传成功
	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}
