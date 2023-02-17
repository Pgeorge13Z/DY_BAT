package main

import (
	"DY_BAT/cmd/user/dal"
	"DY_BAT/cmd/user/kitex_gen/user/userservice"
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
	//addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	//
	//dal.Init()
	//svr := userservice.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))
	////svr := user.NewServer(new(UserServiceImpl))
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

	dal.Init()

	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
		server.WithServiceAddr(&net.TCPAddr{Port: consts.UserServicePort}),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 200}),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
