package rpc

import (
	"DY_BAT/cmd/publish/kitex_gen/publish"
	"DY_BAT/cmd/publish/kitex_gen/publish/publishservice"
	"DY_BAT/pkg/consts"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var PublishClient publishservice.Client

func InitPublish() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		fmt.Println(err)
	}

	//client, err := publishservice.NewClient(consts.PublishServiceName, client2.WithResolver(r))
	client, err := publishservice.NewClient(
		consts.PublishServiceName,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.PublishServiceName}),
		client.WithResolver(r),
		client.WithMuxConnection(1),
	)

	if err != nil {
		log.Fatalln(err)
	}
	PublishClient = client
}

func PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (r *publish.DouyinPublishActionResponse, err error) {
	resp, err := PublishClient.PublishAction(ctx, req)
	if err != nil {
		msg := "PublishAction fail"
		fmt.Println(err)
		return &publish.DouyinPublishActionResponse{
			StatusMsg:  &msg,
			StatusCode: 1,
		}, err
	}
	return resp, err
}

func PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (r *publish.DouyinPublishListResponse, err error) {
	resp, err := PublishClient.PublishList(ctx, req)
	if err != nil {
		msg := "PublishList fail"
		return &publish.DouyinPublishListResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
			VideoList:  make([]*publish.Video, 0),
		}, err
	}
	return resp, err
}
