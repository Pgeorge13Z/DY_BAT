package dal

import (
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
	AddVideo(video *Video) error
	FindByPlayUrl(PlayUrl string) (*Video, error)
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

func (v *VideoImp) AddVideo(video *Video) error {
	if err := v.db.Create(video).Error; err != nil {
		return err
	}
	return nil
}

func (v *VideoImp) FindByPlayUrl(PlayUrl string) (*Video, error) {
	var video Video
	if err := v.db.Where("play_url = ?", PlayUrl).First(&video).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &video, nil
}
