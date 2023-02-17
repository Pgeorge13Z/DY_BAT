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
			UserId:   0,
			Token:    "",
			BaseResp: &user.BaseResp{StatusCode: 1, StatsuMsg: &msg},
		}, err
	}

	return resp, err

}
