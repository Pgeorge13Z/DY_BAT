package dal

import (
	sqlscript "DY_BAT/sql/script"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

var (
	videoService     VideoService
	videoServiceOnce sync.Once
)

type VideoService interface {
	PublishVideo(video *Video) error
}

type VideoServiceImp struct {
	db *gorm.DB
}

func GetVideoService() VideoService {
	videoServiceOnce.Do(func() {
		videoService = &VideoServiceImp{
			db: sqlscript.GetDB(),
		}
	})
	return videoService
}

func (v *VideoServiceImp) PublishVideo(video *Video) error {
	_, err := GetVideoDao().FindByPlayUrl(video.CoverUrl)
	if err == nil {
		fmt.Println("video is already exist , don't Publish again")
		return err
		//return errors.New("video is already exist , don't Publish again")
	}

	if err = GetVideoDao().AddVideo(video); err != nil {
		return err
	}
	return err
}
