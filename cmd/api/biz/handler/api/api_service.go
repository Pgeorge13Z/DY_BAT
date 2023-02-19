// Code generated by hertz generator.

package api

import (
	api "DY_BAT/cmd/api/biz/model/api"
	"DY_BAT/cmd/api/biz/rpc"
	"DY_BAT/cmd/user/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var msg string

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := &api.DouyinUserRegisterResponse{BaseResp: &api.BaseResp{StatsuMsg: &msg}}
	//username := c.Query("username")
	username := req.Username
	//password := c.Query("password")
	password := req.Password
	//
	userResp, _ := rpc.RegisterUser(ctx, &user.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})
	resp.BaseResp.StatsuMsg = userResp.BaseResp.StatsuMsg
	resp.BaseResp.StatusCode = userResp.BaseResp.StatusCode
	resp.Token = userResp.Token
	resp.UserID = userResp.UserId
	////
	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := &api.DouyinUserLoginResponse{BaseResp: &api.BaseResp{StatsuMsg: &msg}}
	username := c.Query("username")
	password := c.Query("password")
	userResp, _ := rpc.LoginUser(ctx, &user.DouyinUserLoginRequest{
		Username: username,
		Password: password,
	})
	resp.BaseResp.StatsuMsg = userResp.BaseResp.StatsuMsg
	resp.BaseResp.StatusCode = userResp.BaseResp.StatusCode
	resp.Token = userResp.Token
	resp.UserID = userResp.UserId
	c.JSON(consts.StatusOK, resp)
}

// UserInfo .
// @router /douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//resp := new(api.DouyinUserResponse)
	resp := &api.DouyinUserResponse{BaseResp: &api.BaseResp{StatsuMsg: &msg}, User: api.NewUser()}
	//Userid := c.Query("user_id")
	Userid := req.GetUserID()
	UserToken := req.GetToken()
	userResp, err := rpc.GetUserInfo(ctx, &user.DouyinUserRequest{UserId: Userid, Token: UserToken})
	resp.BaseResp.StatsuMsg = userResp.BaseResp.StatsuMsg
	resp.BaseResp.StatusCode = userResp.BaseResp.StatusCode
	resp.User.Name = userResp.User.Name
	resp.User.ID = userResp.User.Id
	resp.User.FollowerCount = userResp.User.FollowerCount
	resp.User.FollowCount = userResp.User.FollowCount
	resp.User.IsFollow = userResp.User.IsFollow

	c.JSON(consts.StatusOK, resp)
}
