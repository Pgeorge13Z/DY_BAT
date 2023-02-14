package main

import (
	"DY_BAT/cmd/follow/dao"
	follow "DY_BAT/cmd/follow/kitex_gen/follow"
	"DY_BAT/cmd/user/dal/db_mysql"
	"DY_BAT/pkg/tools"
	"context"
	"sync"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// FollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.DouyinFollowListRequest) (resp *follow.DouyinFollowListResponse, err error) {
	//jwt核验

	if _, err := tools.ParseToken(req.Token); err != nil {
		return nil, err
	}

	//调用数据库层函数，返回关注的id数组
	ids, err := dao.FindFollowIds(req.UserId)
	rsp := follow.NewDouyinFollowListResponse()
	if nil != err {
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		rsp.UserList = nil
		return rsp, nil
	}
	//id为空，返回
	if nil == ids {
		rsp.StatusCode = 0
		rsp.StatusMsg = ""
		rsp.UserList = nil
	}
	//通过id数组获取对应的用户信息，构建用户信息数组
	size := len(ids)
	//等待组
	var wg sync.WaitGroup
	wg.Add(size)
	users := make([]*follow.User, size)
	i, j := 0, 0
	for ; i < size; j++ {
		go func(i int, idx int64) {
			defer wg.Done()
			us, _ := db_mysql.GetUserService().GetUserById(ids[i], req.UserId)
			users[i] = &follow.User{
				Id:            us.Id,
				Name:          us.Name,
				FollowCount:   *us.FollowCount,
				FollowerCount: *us.FollowerCount,
				IsFollow:      us.IsFollow,
			}
		}(i, ids[i])
		i++
	}
	wg.Wait()
	rsp.StatusCode = 0
	rsp.StatusMsg = ""
	rsp.UserList = users
	return rsp, nil
}

// FollowAction implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowAction(ctx context.Context, req *follow.DouyinFollowActionRequest) (resp *follow.DouyinFollowActionResponse, err error) {
	//JWT解析出usrid
	CustomClaim, err := tools.ParseToken(req.Token)
	userId := CustomClaim.User_id
	if err != nil {
		return nil, err
	}
	targetId := req.ToUserId
	//查询数据库表中是否有两者关注记录
	relation, err := dao.FindRelation(userId, targetId)
	rsp := follow.NewDouyinFollowActionResponse()
	//如果记录不为空，更新记录即可
	if relation != nil {
		_, err := dao.UpdateFollowRelation(userId, targetId, 1)
		if err != nil {
			rsp.StatusCode = 1
			rsp.StatusMsg = err.Error()
			return rsp, nil
		}
		rsp.StatusCode = 0
		rsp.StatusMsg = "follow successfully"
		return rsp, nil
	}
	//如果记录为空，插入关注关系
	if _, err := dao.InserFollowRelation(userId, targetId); err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		return rsp, nil
	}
	resp.StatusCode = 0
	resp.StatusMsg = "follow successfully"
	return rsp, nil
}
