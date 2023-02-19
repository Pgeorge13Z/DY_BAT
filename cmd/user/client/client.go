//package main
//
//import (
//	"DY_BAT/cmd/user/kitex_gen/user"
//	"DY_BAT/cmd/user/kitex_gen/user/userservice"
//	"context"
//	client2 "github.com/cloudwego/kitex/client"
//	"log"
//)
//
//func main() {
//
//	client, err := userservice.NewClient("userservice", client2.WithHostPorts("127.0.0.1:9999"))
//	if err != nil {
//		log.Fatalln(err)
//	}
//	UserRegisterRequest := user.NewDouyinUserRegisterRequest()
//	UserRegisterRequest.Username = "zxj2"
//	UserRegisterRequest.Password = "1412"
//	resp, err := client.UserRegister(context.Background(), UserRegisterRequest)
//	if err != nil {
//		log.Fatalln("client userRegister fail")
//		log.Fatalln(err)
//		return
//	}
//
//	log.Println(*resp.BaseResp.StatsuMsg)
//
//}

package main

import (
	"DY_BAT/cmd/user/kitex_gen/user"
	"DY_BAT/cmd/user/kitex_gen/user/userservice"
	"DY_BAT/pkg/consts"
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

	client, err := userservice.NewClient(consts.UserServiceName, client2.WithResolver(r))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(client)
	//UserRegisterRequest := user.NewDouyinUserRegisterRequest()
	//UserRegisterRequest.Username = "zxj6"
	//UserRegisterRequest.Password = "1412"
	//resp, err := client.UserRegister(context.Background(), UserRegisterRequest)
	//if err != nil {
	//	log.Fatalln("client userRegister fail")
	//	log.Fatalln(err)
	//	return
	//}
	//
	//log.Println(*resp.BaseResp.StatsuMsg)

	//LoginRequest := user.NewDouyinUserLoginRequest()
	//LoginRequest.Username = "zxj"
	//LoginRequest.Password = "123"
	//resp, err := client.UserLogin(context.Background(), LoginRequest)
	//if err != nil {
	//	log.Fatalln("client userlogin fail")
	//	log.Fatalln(err)
	//	return
	//}
	//
	//fmt.Println(*resp.BaseResp.StatsuMsg)

	UserRequest := user.NewDouyinUserRequest()
	UserRequest.UserId = 7
	UserRequest.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Inp4ajYiLCJ1c2VyX2lkIjo3LCJpc3MiOiJ6eGoiLCJleHAiOjE2NzY4OTcyMjZ9.csaZd7ADUzyuU26-gHs3TPBi8CLtzvGIfMg74AAHG4Q"
	resp, err := client.UserInfo(context.Background(), UserRequest)
	if err != nil {
		log.Fatalln("client userInfo fail")
		log.Fatalln(err)
		return
	}

	//fmt.Println(*resp.BaseResp.StatsuMsg)
	fmt.Printf("%#v", resp.User)

}
