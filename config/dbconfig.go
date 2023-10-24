package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"plan/model"
)

var DB *gorm.DB

// 用于连接数据库和初始化数据库等等操作
func InitMysql() error {
	var err error
	dsn := "root:abc8909389@tcp(172.17.48.147:3306)/golearn?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		panic(err)
	}
	return migration()
}

// 创建更新对应表
func migration() error {
	//自动映射和创建对应结构体
	return DB.AutoMigrate(new(model.Todo))
}
