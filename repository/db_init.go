package repository

import (
	"Douyin/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {

	info := config.Conf.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		info.User, info.Pass, info.Host, info.Port, info.Dbname,
		info.Charset, info.ParseTime, info.Loc)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动建表
	if err = DB.AutoMigrate(User{}, Video{}, Comment{}, Favorite{}, Relation{}, Message{}); err != nil {
		return err
	}
	return nil
}
