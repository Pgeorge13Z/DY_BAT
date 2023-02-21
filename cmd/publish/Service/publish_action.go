package Service

import (
	"DY_BAT/cmd/publish/dal"
	"DY_BAT/cmd/publish/kitex_gen/publish"
	"DY_BAT/pkg/tools"
	"context"
	"fmt"
	"github.com/godruoyi/go-snowflake"
	"strconv"
	"strings"
	"time"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{ctx: ctx}
}

func (s *PublishService) PublishAction(req *publish.DouyinPublishActionRequest) error {
	Token := req.Token
	Title := req.Title
	Data := req.Data
	Claims, err := tools.ParseToken(Token)
	if err != nil {
		fmt.Println("parsetoken err:", err)
		return err
	}
	UserID := Claims.User_id

	SnowflakeId := snowflake.ID()
	filename := strings.Join([]string{strconv.Itoa(int(SnowflakeId)), ".mp4"}, "")

	filePath := strings.Join([]string{tools.GetPath(), "/public/", filename}, "")

	//上传视频到本地,后续可更改为上传到OSS
	err = tools.PublishVideoToPublic(Data, filePath)
	if err != nil {
		fmt.Println("PublishVideoToPublic fail ")
		return err
	}

	//获取视频的播放地址，这里用视频本地地址代替，后续可以采用上传到OSS改进
	videoUrl := filePath

	//获取视频封面，这里用固定图片代替，后续可以采用ffmpeg取图
	PicName := "cover.png"
	PicPath := strings.Join([]string{tools.GetPath(), "/public/", PicName}, "")

	//publish video
	video := &dal.Video{
		UserId:        UserID,
		PlayUrl:       videoUrl,
		CoverUrl:      PicPath,
		Title:         Title,
		FavoriteCount: 0,
		CommentCount:  0,
		Time:          time.Now().Unix(),
	}
	err = dal.GetVideoService().PublishVideo(video)
	if err != nil {
		fmt.Println("PublishVideo fail: ", err)
		return err
	}
	return err
}
