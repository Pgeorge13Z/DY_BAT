package main

import (
	"DY_BAT/cmd/user/dal"
	"DY_BAT/cmd/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	dal.Init()
	svr := userservice.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))
	//svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
