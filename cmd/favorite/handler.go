package main

import (
	"DY_BAT/cmd/favorite/dal/db"
	favorite "DY_BAT/cmd/favorite/kitex_gen/favorite"
	"DY_BAT/cmd/favorite/kitex_gen/feed"
	"DY_BAT/pkg/tools"
	"context"
	"github.com/cloudwego/hertz/pkg/common/json"
	"log"
	"sync"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	// TODO: Your code here...
	var msg string
	resp = &favorite.DouyinFavoriteActionResponse{StatusMsg: &msg}
	CustomClaim, err := tools.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}

	user_id := CustomClaim.User_id
	video_id := req.VideoId
	action_type := req.ActionType
	// select the record from the like table
	likeRecord, err := db.SelectFavoriteRecord(user_id, video_id) // err is equal nil-->exist the record

	if action_type == 1 { // favorite action
		if err != nil { // record is nil--> insert   tested
			isSucessInsert, _ := db.InsertFavoriteRecord(user_id, video_id, action_type)
			if isSucessInsert == false {
				resp.StatusCode = 1
				*resp.StatusMsg = "Failed to favourite"
				log.Println(*resp.StatusMsg)
				//*resp.BaseResp.StatusMsg = "Failed to favourite"
				return resp, nil
			}
			resp.StatusCode = 0
			*resp.StatusMsg = "Favourite successfully!"
			return resp, nil
		} else {
			// update
			isSuccessUpdate, _ := db.UpdateFavoriteRecord(user_id, video_id, action_type)
			if isSuccessUpdate == false {
				resp.StatusCode = 1
				*resp.StatusMsg = "Failed to update!"
				return resp, nil
			}
			resp.StatusCode = 0
			*resp.StatusMsg = "Favourite successfully!"
			return resp, nil
		}
	} else if action_type == 2 {
		//Cancel likes
		if err != nil {
			resp.StatusCode = 1
			*resp.StatusMsg = "Failed to cancel favourite"
			return resp, nil
		}
		if likeRecord.IsFavorite == 1 {
			// update
			_, err := db.UpdateFavoriteRecord(user_id, video_id, action_type)
			if err != nil {
				resp.StatusCode = 1
				*resp.StatusMsg = "Failed!"
				return resp, nil
			}
			resp.StatusCode = 0
			*resp.StatusMsg = "Cancle favourtie successfully!"
			return resp, nil
		}
		resp.StatusCode = 1
		*resp.StatusMsg = "Failed to cancel favourite"
		return resp, nil
	}
	// other action
	resp.StatusCode = 1
	*resp.StatusMsg = "Illegal operation!"
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) { // tested
	// TODO: Your code here...
	var msg string
	resp = &favorite.DouyinFavoriteListResponse{StatusMsg: msg}
	//jwt核验
	if _, err := tools.ParseToken(req.Token); err != nil {
		return nil, err
	}

	//调用数据库层函数，返回like的id数组
	ids, err := db.SelectLikeVideoIds(req.UserId)
	log.Println(err)
	//	rsp := follow.NewDouyinFollowListResponse()
	if nil != err {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		resp.VideoList = nil
		return resp, nil
	}
	//id为空，返回

	if nil == ids {
		resp.StatusCode = 0
		resp.StatusMsg = "like list is null"
		resp.VideoList = nil
		return resp, nil
	}
	//通过id数组获取对应的video
	size := len(ids)
	log.Println(size)
	//等待组
	var wg sync.WaitGroup
	wg.Add(size) // size task
	videos := make([]*feed.Video, size)
	i := 0
	for ; i < size; i++ {
		go func(i int, idx int64) {
			defer wg.Done()
			vs, _ := db.SelectVideobyId(idx)
			str := vs.Author
			var p db.Users
			json.Unmarshal([]byte(str), &p)
			//{"Name": "Platypus", "Order": "Monotremata"}
			author := &feed.User{
				Id:              p.UserId,
				Name:            p.Username,
				FollowCount:     &p.FollowCount,
				FollowerCount:   &p.FollowerCount,
				IsFollow:        true, // TODO
				Avatar:          &p.Avatar,
				BackgroundImage: &p.BackgroundImage,
				Signature:       &p.Signature,
				TotaolFavorited: &p.TotalFavorite,
				WorkCount:       &p.WorkCount,
				FavoriteCount:   &p.FavoriteCount,
			}
			videos[i] = new(feed.Video)
			videos[i].Author = new(feed.User)
			videos[i] = &feed.Video{
				Id:            vs.VideoId,
				Author:        author,
				PlayUrl:       vs.PlayUrl,
				CoverUrl:      vs.CoverUrl,
				FavoriteCount: vs.FavoriteCount,
				CommentCount:  vs.CommentCount,
				IsFavorite:    vs.IsFavorite,
				Title:         vs.Title,
			}
		}(i, ids[i])
	}
	wg.Wait()
	resp.StatusCode = 0
	resp.StatusMsg = "Get likelist successfully!"
	resp.VideoList = videos
	return resp, nil
}
