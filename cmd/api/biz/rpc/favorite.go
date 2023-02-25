package rpc

import (
	"DY_BAT/cmd/favorite/kitex_gen/favorite"
	"DY_BAT/cmd/favorite/kitex_gen/favorite/favoriteservice"
	"DY_BAT/cmd/favorite/kitex_gen/feed"
	"DY_BAT/pkg/consts"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var favoriteclient favoriteservice.Client

func InitFavorite() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		log.Fatalln(err)
	}

	client, err := favoriteservice.NewClient(
		consts.FavoriteServiceName,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.FavoriteServiceName}),
		client.WithResolver(r),
		client.WithMuxConnection(1),
	)
	if err != nil {
		log.Fatalln(err)
	}
	favoriteclient = client
}
func ActionFavorite(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (r *favorite.DouyinFavoriteActionResponse, err error) {
	resp, err := favoriteclient.FavoriteAction(ctx, req)
	if err != nil {
		msg := "favoriteClient Action fail"
		return &favorite.DouyinFavoriteActionResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
		}, err
	}
	return resp, err
}

func ListFavorite(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (r *favorite.DouyinFavoriteListResponse, err error) {
	resp, err := favoriteclient.FavoriteList(ctx, req)
	if err != nil {
		msg := "favoriteClient checklist fail"
		videos := make([]*feed.Video, 0)
		return &favorite.DouyinFavoriteListResponse{
			StatusCode: 1,
			StatusMsg:  msg,
			VideoList:  videos,
		}, err
	}
	return resp, err
}
