package main

import (
	"Douyin/config"
	"Douyin/middleware/redis"
	"Douyin/repository"
	"Douyin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := initial(); err != nil {
		panic(err)
	}

	r := gin.Default()
	router.Init(r)
	err := r.Run() // listen and serve on localhost:8080
	if err != nil {
		panic(err)
	}
}

// 根据配置文件初始化数据库和redis
func initial() error {
	go RunMessageServer() // 这一行是demo里自带的，可能写message的部分会用到

	if err := config.Init(); err != nil {
		return err
	}
	if err := repository.Init(); err != nil {
		return err
	}
	redis.Init()

	return nil
}
