package main

import (
	"DY_BAT/cmd/user/dal/db_mysql"
	constants "DY_BAT/pkg/consts"
	sqlscript "DY_BAT/sql/script"
)

func main() {
	sqlscript.InitDB(constants.MySQLDefaultDSN)
	db := sqlscript.GetDB()
	db.AutoMigrate(&db_mysql.User{})

	//创建
	//user := db_mysql.User{
	//	UserId:        2,
	//	Username:      "test2",
	//	Password:      "123",
	//	Avatar:        "test",
	//	Salt:          "test",
	//	FollowCount:   10,
	//	FollowerCount: 100,
	//}
	//
	//db.Create(user)
	//db.Model(&db_mysql.User{}).Create(user)

	//查询

	//var user db_mysql.User
	////user := new(db_mysql.User)
	//db.First(&user)
	//fmt.Printf("%#v", user)
	//
	//var users []db_mysql.User
	//db.Find(&users)
	//fmt.Printf("%#v", users)
	//db.Delete(&user)
}
