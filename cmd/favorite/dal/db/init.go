package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	UserId          int64  `gorm:"column:user_id;PRIMARY_KEY"`
	Username        string `gorm:"column:username;UNIQUE"`
	Password        string `gorm:"column:password;NOT NULL"`
	Avatar          string `gorm:"column:avatar"`
	Salt            string `gorm:"column:Salt"`
	FollowCount     int64  `gorm:"column:follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count"`
	BackgroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:Signature"`
	TotalFavorite   int64  `gorm:"column:TotalFavorite"`
	WorkCount       int64  `gorm:"column:workCount"`
	FavoriteCount   int64  `gorm:"column:favoriteCount"`
}

type Videos struct {
	//Id            int64  `gorm:"primary_key;AUTO_INCREMENT"`
	VideoId       int64  `gorm:"column:video_id;primary_key;"`
	Author        string `gorm:"column:author"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	IsFavorite    bool   `gorm:"column:is_favorite"`
	Title         string `gorm:"column:title"`
}

type Favorites struct { // must start at upper-case letters
	Id         int64 `gorm:"primary_key;AUTO_INCREMENT"`
	UserId     int64 `gorm:"index:uid"`
	VideoId    int64 `gorm:"index:fid"`
	IsFavorite int32
}

// TableName 设置Like结构体对应数据库表名。
func (Users) TableName() string {
	return "Users"
}
func (Videos) TableName() string {
	return "Videos"
}

func (Favorites) TableName() string {
	return "Favorites"
}

var db *gorm.DB

//func GetDB() *gorm.DB {
//	Init()
//	return db
//}

func Init() {
	var err error
	dsn := "root:123456@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	open, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			//预编译sql语句，提高效率（缓存）
			PrepareStmt: true,
			//跳过默认事务
			//如果使用gorm的hook或者关联创建时，false
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := open.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	db = open

	// 迁移 schema
	err = db.AutoMigrate(&Users{})
	if err != nil {
		panic("failed to connect database")
		return
	}
	db.AutoMigrate(&Videos{})
	db.AutoMigrate(&Favorites{})

}
