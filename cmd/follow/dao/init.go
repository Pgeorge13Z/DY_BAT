package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Follow 用户关系表结构
type Follow struct {
	Id         int64 `gorm:"primary_key;AUTO_INCREMENT"`
	UserId     int64 `gorm:"index:uid"`
	FollowerId int64 `gorm:"index:fid"`
	Cancel     int8
}

// TableName 设置Follow结构体对应数据库表名。
func (Follow) TableName() string {
	return "follows"
}

var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
	println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Follow{})
}
