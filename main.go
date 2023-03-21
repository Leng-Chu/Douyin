package main

import (
	"Douyin/config"
	"Douyin/middleware/redis"
	"Douyin/repository"
	"Douyin/router"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	if err := initial(); err != nil {
		panic(err)
	}
	r := gin.Default()
	router.Init(r)
	err := r.Run(":" + strconv.Itoa(config.Conf.Server.Port))
	if err != nil {
		panic(err)
	}
}

// 根据配置文件初始化数据库和redis
func initial() error {
	go RunMessageServer()

	if err := config.Init(); err != nil {
		return err
	}
	if err := repository.Init(); err != nil {
		return err
	}
	redis.Init()

	return nil
}
