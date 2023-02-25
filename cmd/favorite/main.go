package main

import (
	"DY_BAT/cmd/favorite/dal"
	"DY_BAT/cmd/favorite/kitex_gen/favorite/favoriteservice"
	"DY_BAT/pkg/consts"
	"fmt"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	dal.Init()
	//svr := favorite.NewServer(new(FavoriteServiceImpl))
	//
	//err := svr.Run()
	//
	//if err != nil {
	//	log.Println(err.Error())
	//}

	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddress})
	if err != nil {
		fmt.Println(err)
	}

	svr := favoriteservice.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.FavoriteServiceName}),
		server.WithServiceAddr(&net.TCPAddr{Port: consts.FavoriteServicePort}),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 200}),
		//server.WithMuxTransport(),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
