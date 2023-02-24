package main

import (
	"DY_BAT/cmd/publish/dal"
	"DY_BAT/cmd/publish/kitex_gen/publish/publishservice"
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
	//svr := publish.NewServer(new(PublishServiceImpl))
	//
	//err := svr.Run()
	//
	//if err != nil {
	//	log.Println(err.Error())
	//}
	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddress})
	if err != nil {
		fmt.Println("EtcdRegistry err", err)
	}

	dal.Init()

	svr := publishservice.NewServer(new(PublishServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.PublishServiceName}),
		server.WithServiceAddr(&net.TCPAddr{Port: consts.PublishServicePort}),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 200}),
		//server.WithMuxTransport(),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}

}
