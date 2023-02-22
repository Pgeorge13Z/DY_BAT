package dal

import (
	"DY_BAT/cmd/publish/kitex_gen/publish"
	sqlscript "DY_BAT/sql/script"
	"errors"
	"gorm.io/gorm"
	"sync"
)

var (
	videoDao     VideoDao
	videoDaoOnce sync.Once
)

type VideoDao interface {
	AddVideo(video *publish.Video) error
	FindByPlayUrl(PlayUrl string) (*publish.Video, error)
	QueryVideoByUserId(userId int64) ([]*publish.Video, error)
}

type VideoImp struct {
	db *gorm.DB
}

func GetVideoDao() VideoDao {
	videoDaoOnce.Do(func() {
		videoDao = &VideoImp{
			db: sqlscript.GetDB(),
		}
	})
	return videoDao
}

func (v *VideoImp) AddVideo(video *publish.Video) error {
	if err := v.db.Create(video).Error; err != nil {
		return err
	}
	return nil
}

func (v *VideoImp) FindByPlayUrl(PlayUrl string) (*publish.Video, error) {
	var video publish.Video
	if err := v.db.Where("play_url = ?", PlayUrl).First(&video).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &video, nil
}

func (v *VideoImp) QueryVideoByUserId(userId int64) ([]*publish.Video, error) {
	videoList := make([]*publish.Video, 0)
	if err := v.db.Where("user_id=?", userId).Find(&videoList).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return videoList, nil
}
