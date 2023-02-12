package main

import (
	"DY_BAT/cmd/user/dal/db_mysql"
	user "DY_BAT/kitex_gen/user"
	"DY_BAT/pkg/tools"
	"context"
	"sync/atomic"
)

const (
	success = 0
	fail    = 1
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...

	rsp := user.NewDouyinUserRegisterResponse()

	username := req.GetUsername()
	password := req.GetPassword()
	//随机生成salt
	salt := tools.RandomStringUtil()
	//密码加密,内部实现改成了用bcrypt加密
	password = tools.Md5Util(password, salt)

	//更新用户ID
	userIdSequence := db_mysql.GetUserService().FindLastUserId()

	//注册用户
	err := db_mysql.GetUserService().UserRegister(username, password, salt)

	if err != nil {
		resp.UserId = 0
		msg := err.Error()
		resp.BaseResp.StatsuMsg = &msg
		resp.BaseResp.StatusCode = fail

	} else {
		msg := "UserRegister successfully"
		atomic.AddInt64(&userIdSequence, 1)

	}

}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryCurUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryCurUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}
