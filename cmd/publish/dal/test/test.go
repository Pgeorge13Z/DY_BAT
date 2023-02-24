package main

import (
	"DY_BAT/cmd/publish/dal"
	"fmt"
)

func main() {
	dal.Init()
	//video := &dal.Video{
	//	UserId:        1,
	//	PlayUrl:       "test",
	//	CoverUrl:      "test",
	//	Title:         "test",
	//	FavoriteCount: 0,
	//	CommentCount:  0,
	//	Time:          time.Now().Unix(),
	//}
	//dal.GetVideoService().PublishVideo(video)
	videos := dal.GetVideoService().PublishList(1)
	fmt.Printf("%T\n", videos)
	for _, video := range videos {
		fmt.Printf("%T\n", video)
		fmt.Println(video)
	}
}
