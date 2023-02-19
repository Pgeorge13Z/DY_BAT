package rpc

import (
	"DY_BAT/cmd/user/kitex_gen/user"
	"DY_BAT/cmd/user/kitex_gen/user/userservice"
	"DY_BAT/pkg/consts"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var userclient userservice.Client

func InitUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		log.Fatalln(err)
	}

	client, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
		client.WithResolver(r),
		client.WithMuxConnection(1),
	)
	if err != nil {
		log.Fatalln(err)
	}
	userclient = client
}

func RegisterUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (r *user.DouyinUserRegisterResponse, err error) {
	resp, err := userclient.UserRegister(ctx, req)

	if err != nil {
		msg := "userClient Register fail"
		return &user.DouyinUserRegisterResponse{
			UserId:     0,
			Token:      "",
			StatusCode: 1,
			StatusMsg:  &msg,
		}, err
	}

	return resp, err

}

func LoginUser(ctx context.Context, req *user.DouyinUserLoginRequest) (r *user.DouyinUserLoginResponse, err error) {
	resp, err := userclient.UserLogin(ctx, req)
	if err != nil {
		msg := "userClient Login fail"
		return &user.DouyinUserLoginResponse{
			UserId:     0,
			Token:      "",
			StatusCode: 1,
			StatusMsg:  &msg,
		}, err
	}
	return resp, err
}

func GetUserInfo(ctx context.Context, req *user.DouyinUserRequest) (r *user.DouyinUserResponse, err error) {
	resp, err := userclient.UserInfo(ctx, req)
	if err != nil {
		msg := "userClient GetInfo fail"
		return &user.DouyinUserResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
			User:       user.NewUser(),
		}, err
	}
	return resp, err
}
