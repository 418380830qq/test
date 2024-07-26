package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "root:123456@tcp(127.0.0.1:3308)/"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库: " + err.Error())
	}

	// 创建数据库
	err = db.Exec("CREATE DATABASE IF NOT EXISTS kainengpaichong").Error
	if err != nil {
		panic("无法创建数据库: " + err.Error())
	}

	// 连接到数据库
	db, err = gorm.Open(mysql.Open(dsn+"kainengpaichong?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("无法连接到新创建的数据库: " + err.Error())
	}

	//进行数据库迁移
	if err := db.AutoMigrate(&Fromplan{}); err != nil {
		panic("数据库迁移失败: " + err.Error())
	}

	DB = db
}
