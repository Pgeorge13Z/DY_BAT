package main

import (
	"DY_BAT/cmd/publish/Service"
	publish "DY_BAT/cmd/publish/kitex_gen/publish"
	"context"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	var msg string
	resp = &publish.DouyinPublishActionResponse{StatsuMsg: &msg}

	err = Service.NewPublishService(ctx).PublishAction(req)

	if err != nil {
		msg = "视频投稿失败"
		resp.StatusCode = 1
		resp.StatsuMsg = &msg
		return resp, err
	}
	msg = "视频投稿成功"
	resp.StatusCode = 0
	resp.StatsuMsg = &msg
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	var msg string
	resp = &publish.DouyinPublishListResponse{VideoList: make([]*publish.Video, 0), StatsuMsg: &msg}

	resp.VideoList, err = Service.NewPublishService(ctx).PublishList(req)

	return
}
