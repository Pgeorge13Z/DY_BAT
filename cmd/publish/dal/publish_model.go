package dal

const TableNameUser string = "video"

type Video struct {
	UserId        int64  `json:"user_id" gorm:"index,unique;not null"`
	PlayUrl       string `json:"play_url" gorm:"not null"`
	CoverUrl      string `json:"cover_url" gorm:"not null"`
	Title         string `json:"title" gorm:"not null"`
	FavoriteCount int64  `gorm:"default:0"`
	CommentCount  int64  `gorm:"default:0"`
	Time          int64  `gorm:"not null"`
}

func (u *Video) TableName() string {
	return TableNameUser
}
