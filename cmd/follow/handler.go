package main

import (
	"context"
	follow "followservice/kitex_gen/follow"
	"dao"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// FollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.DouyinFollowListRequest) (resp *follow.DouyinFollowListResponse, err error) {
	//jwt核验

	/*_, err := tools.ParseToken(req.Token).User_id
	if (err != nil) {
		return nil err
	}*/
	
	//调用数据库层函数，返回关注的id数组
	ids, err := dao.FindFollowIds(req.UserID)
	resp := NewDouyinFollowListResponse()
	if (nil != err) {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		resp.UserList = nil
		return resp, nil
	}
	//id为空，返回
	if (nil == ids) {
		resp.StatusCode = 0
		resp.StatusMsg = nil
		resp.UserList = nil
	}
	//通过id数组获取对应的用户信息，构建用户信息数组
	size := len(ids)
	//等待组
	var wg sync.WaitGroup
	wg.Add(len)
	users := make([]User, size)
	i, j := 0, 0
	for ; i < len; j++ {
		go func(i int, idx int64) {
			defer wg.Done()
			users[i], _ = f.GetUserByIdWithCurId(idx, userId)
		}(i, ids[i])
		i++
	}
	wg.Wait()
	resp.StatusCode = 0
	resp.StatusMsg = nil
	resp.UserList = users
	return resp, nil
}

// FollowAction implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowAction(ctx context.Context, req *follow.DouyinFollowActionRequest) (resp *follow.DouyinFollowActionResponse, err error) {
	//JWT解析出usrid
	userId, err := tools.ParseToken(req.Token).User_id
	if (err != nil) {
		return nil err
	}
	targetId := req.ToUserId
	//查询数据库表中是否有两者关注记录
	relation, err := dao.FindRelation(userId, targetId)
	resp := NewDouyinFollowActionResponse()
	//如果记录不为空，更新记录即可
	if (relation != nil) {
		re, err := dao.UpdateFollowRelation(userId, targetId, 1)
		if (err != nil) {
			resp.StatusCode = 1
			resp.StatusMsg = err.Error()
			return resp
		}
		resp.StatusCode = 0
		resp.StatusMsg = "follow successfully"
		return resp
	}
	//如果记录为空，插入关注关系
	re, err := dao.InserFollowRelation(userId, targetId)
	if (err != nil) {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		return resp
	}
	resp.StatusCode = 0
	resp.StatusMsg = "follow successfully"
	return resp,nil
}
