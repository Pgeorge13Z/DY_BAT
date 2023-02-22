package dal

import (
	"DY_BAT/cmd/publish/kitex_gen/publish"
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
	PublishVideo(video *publish.Video) error
	PublishList(userId int64) []*publish.Video
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

func (v *VideoServiceImp) PublishVideo(video *publish.Video) error {
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

func (v *VideoServiceImp) PublishList(userId int64) []*publish.Video {
	VideoList, err := GetVideoDao().QueryVideoByUserId(userId)
	if err != nil {
		fmt.Println("PublishList fail:", err)
		return make([]*publish.Video, 0)
	}
	return VideoList
}
