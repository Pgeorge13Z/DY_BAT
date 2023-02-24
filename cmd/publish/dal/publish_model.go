package dal

//import (
//	"DY_BAT/cmd/user/kitex_gen/user"
//	"gorm.io/gorm"
//)
//
//const TableNameUser string = "video"
//
////type Video struct {
////	UserId        int64  `json:"user_id" gorm:"index,unique;not null"`
////	PlayUrl       string `json:"play_url" gorm:"not null"`
////	CoverUrl      string `json:"cover_url" gorm:"not null"`
////	Title         string `json:"title" gorm:"not null"`
////	FavoriteCount int64  `gorm:"default:0"`
////	CommentCount  int64  `gorm:"default:0"`
////	Time          int64  `gorm:"not null"`
////}
//
//type Video struct {
//	gorm.Model
//	Id            int64      `thrift:"id,1,required" frugal:"1,required,i64" json:"id" gorm:"index,unique;not null;PRIMARY_KEY"`
//	Author        *user.User `thrift:"author,2,required" frugal:"2,required,user.User" json:"author" gorm:"-"`
//	PlayUrl       string     `thrift:"play_url,3,required" frugal:"3,required,string" json:"play_url" gorm:"not null"`
//	CoverUrl      string     `thrift:"cover_url,4,required" frugal:"4,required,string" json:"cover_url" gorm:"not null"`
//	FavoriteCount int64      `thrift:"favorite_count,5,required" frugal:"5,required,i64" json:"favorite_count" gorm:"default:0"`
//	CommentCount  int64      `thrift:"comment_count,6,required" frugal:"6,required,i64" json:"comment_count" gorm:"default:0"`
//	IsFavorite    bool       `thrift:"is_favorite,7,required" frugal:"7,required,bool" json:"is_favorite" gorm:"not null"`
//	Title         string     `thrift:"title,8,required" frugal:"8,required,string" json:"title" gorm:"not null"`
//}
//
//func (u *Video) TableName() string {
//	return TableNameUser
//}
