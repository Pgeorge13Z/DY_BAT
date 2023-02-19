package main

import (
	"DY_BAT/cmd/user/dal/db_mysql"
	user "DY_BAT/cmd/user/kitex_gen/user"
	"DY_BAT/pkg/tools"
	"context"
	"sync/atomic"
)

const (
	success = 0
	fail    = 0
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	var msg string
	resp = &user.DouyinUserRegisterResponse{BaseResp: &user.BaseResp{StatsuMsg: &msg}}
	username := req.GetUsername()
	password := req.GetPassword()
	//随机生成salt
	salt := tools.RandomStringUtil()
	//密码加密,内部实现改成了用bcrypt加密
	password = tools.Md5Util(password, salt)

	//更新用户ID
	userIdSequence := db_mysql.GetUserService().FindLastUserId()

	//注册用户
	err = db_mysql.GetUserService().UserRegister(username, password, salt)
	//
	if err != nil {
		resp.UserId = 0
		resp.Token = " "
		msg = err.Error()
		resp.BaseResp.StatsuMsg = &msg
		resp.BaseResp.StatusCode = fail

	} else {
		token, err := tools.GenToken(username, userIdSequence)
		if err != nil {
			resp.UserId = 0
			resp.Token = " "
			msg = "token generation failed" + err.Error()
			resp.BaseResp.StatsuMsg = &msg
			resp.BaseResp.StatusCode = fail
		}

		msg = "UserRegister successfully"
		//此处有一个bug,记录个数为0和1的时候，主键都为1，因此第一条记录的id为2，本项目将第一条记录内置，因此默认是从第二条记录开始使用.
		atomic.AddInt64(&userIdSequence, 1)
		resp.UserId = userIdSequence
		resp.Token = token
		resp.BaseResp.StatsuMsg = &msg
		resp.BaseResp.StatusCode = success

	}

	msg = "UserRegister successfully"
	resp.BaseResp.StatsuMsg = &msg
	resp.BaseResp.StatusCode = success
	return resp, err
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	var msg string
	resp = &user.DouyinUserLoginResponse{BaseResp: &user.BaseResp{StatsuMsg: &msg}}
	username := req.GetUsername()
	password := req.GetPassword()
	userResp, err := db_mysql.GetUserService().UserLogin(username, password)
	//
	if err != nil {
		msg = "Success failed"
		resp.BaseResp.StatsuMsg = &msg
		resp.BaseResp.StatusCode = fail
	}
	resp.UserId = userResp.UserId
	token, _ := tools.GenToken(username, userResp.UserId)
	resp.Token = token
	msg = "Success login"
	resp.BaseResp.StatsuMsg = &msg
	resp.BaseResp.StatusCode = success

	return resp, err
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	UserId := req.GetUserId()
	var msg string
	resp = &user.DouyinUserResponse{BaseResp: &user.BaseResp{StatsuMsg: &msg}}

	userRsp, err := db_mysql.GetUserService().GetUserById(UserId)
	if err != nil {
		msg := "Get UserInfo fail"
		resp.BaseResp.StatsuMsg = &msg
		resp.BaseResp.StatusCode = fail
	}
	resp.User = userRsp
	msg = "Get Userinfo success"
	resp.BaseResp.StatsuMsg = &msg
	resp.BaseResp.StatusCode = success

	return resp, err
}
