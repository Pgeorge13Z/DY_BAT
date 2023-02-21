// Code generated by Kitex v0.4.4. DO NOT EDIT.

package feedservice

import (
	feed "DY_BAT/cmd/feed/kitex_gen/feed"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return feedServiceServiceInfo
}

var feedServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FeedService"
	handlerType := (*feed.FeedService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetUserFeed": kitex.NewMethodInfo(getUserFeedHandler, newFeedServiceGetUserFeedArgs, newFeedServiceGetUserFeedResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "feed",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func getUserFeedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*feed.FeedServiceGetUserFeedArgs)
	realResult := result.(*feed.FeedServiceGetUserFeedResult)
	success, err := handler.(feed.FeedService).GetUserFeed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedServiceGetUserFeedArgs() interface{} {
	return feed.NewFeedServiceGetUserFeedArgs()
}

func newFeedServiceGetUserFeedResult() interface{} {
	return feed.NewFeedServiceGetUserFeedResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (r *feed.DouyinFeedResponse, err error) {
	var _args feed.FeedServiceGetUserFeedArgs
	_args.Req = req
	var _result feed.FeedServiceGetUserFeedResult
	if err = p.c.Call(ctx, "GetUserFeed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}