package main

import (
	"DY_BAT/cmd/publish/dal"
	"time"
)

func main() {
	dal.Init()
	video := &dal.Video{
		UserId:        1,
		PlayUrl:       "test",
		CoverUrl:      "test",
		Title:         "test",
		FavoriteCount: 0,
		CommentCount:  0,
		Time:          time.Now().Unix(),
	}
	dal.GetVideoService().PublishVideo(video)
}
