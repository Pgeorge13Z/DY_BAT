package main

import (
	"DY_BAT/cmd/favorite/kitex_gen/favorite"
	"DY_BAT/cmd/favorite/kitex_gen/favorite/favoriteservice"
	"DY_BAT/pkg/consts"
	"DY_BAT/pkg/tools"
	"context"
	"fmt"
	client2 "github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		fmt.Println(err)
	}

	client, err := favoriteservice.NewClient(consts.FavoriteServiceName, client2.WithResolver(r))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(client)

	likeActionRequest := favorite.NewDouyinFavoriteActionRequest()
	// token
	var i1 int64 = 7
	token, err := tools.GenToken("j1", i1)

	likeActionRequest.Token = token
	likeActionRequest.VideoId = 12
	likeActionRequest.ActionType = 2
	log.Println(token)

	resp, err := client.FavoriteAction(context.Background(), likeActionRequest)

	// like list test
	/*
		{"UserId":1, "Username":"jingtao", "FollowCount":3, "FollowerCount":5, Avatar:"test.png", BackgroundImage: "back.png", "Signature": "this is test user...", "TotalFavorite": 1, "WorkCount": 1, favorite_count: 1}
	*/
	//likelistRequest := favorite.NewDouyinFavoriteListRequest()
	//var i1 int64 = 3
	//token, err := tools.GenToken("j1", i1)
	//likelistRequest.UserId = i1
	//likelistRequest.Token = token
	//resp, err := client.FavoriteList(context.Background(), likelistRequest)
	//
	if err != nil {
		log.Fatalln("client favoriteAction fail")
		log.Fatalln(err)
		log.Println(token)
		return
	}
	log.Println(resp.StatusCode, *resp.StatusMsg)
}
