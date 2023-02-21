package main

import (
	"DY_BAT/cmd/comment/dao"
	comment "DY_BAT/cmd/comment/kitex_gen/comment"
	"DY_BAT/cmd/user/dal/db_mysql"
	constants "DY_BAT/pkg/consts"
	"DY_BAT/pkg/tools"
	"context"
	"log"
	"sort"
	"sync"
	"time"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	var msg string
	//jwt核验
	if _, err := tools.ParseToken(req.Token); err != nil {
		return nil, err
	}
	userid, _ := tools.ParseToken(req.Token)
	if req.ActionType == 1 {
		//发表评论数据准备
		var sendComment dao.Comment
		sendComment.UserId = userid.User_id
		sendComment.CommentText = req.GetCommentText()
		sendComment.VideoId = req.VideoId
		sendComment.CreateDate = time.Now()
		//发表评论
		commentInfo, err := s.Send(sendComment)
		if err != nil {
			msg = err.Error()
			resp.BaseResp.StatusCode = 1
			resp.BaseResp.StatusMsg = &msg
			return resp, err
		}
		msg = ""
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMsg = &msg
		resp.Comment = &commentInfo
		return resp, nil
	}
	//删除评论
	if req.ActionType == 2 {
		commentId := req.GetCommentId()
		err2 := s.Delete(commentId)
		if err2 != nil {
			msg = err2.Error()
			resp.BaseResp.StatusCode = 1
			resp.BaseResp.StatusMsg = &msg
			return resp, err2
		}
		return
	}
	return
}

func (s *CommentServiceImpl) Send(cmt dao.Comment) (comment.Comment, error) {
	//发表评论
	var info dao.Comment
	info.VideoId = cmt.VideoId               //评论视频id传入
	info.UserId = cmt.UserId                 //评论用户id传入
	info.CommentText = cmt.CommentText       //评论内容传入
	info.ActionType = constants.ValidComment //评论状态，1，有效
	info.CreateDate = cmt.CreateDate         //评论时间

	//1.评论信息存储：
	commentRtn, err := dao.InsertComment(info)
	if err != nil {
		return comment.Comment{}, err
	}
	//2.查询用户信息
	userData, err2 := db_mysql.GetUserService().GetUserById(cmt.UserId)
	if err2 != nil {
		return comment.Comment{}, err2
	}
	//3.拼接
	commentData := comment.Comment{
		Id:         commentRtn.Id,
		User:       (*comment.User)(userData),
		Content:    commentRtn.CommentText,
		CreateDate: commentRtn.CreateDate.Format(constants.DateTime),
	}
	//缓存
	//返回结果
	return commentData, nil
}

func (s *CommentServiceImpl) Delete(commentId int64) error {
	return dao.DeleteComment(commentId)
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	var msg string
	//jwt核验
	if _, err := tools.ParseToken(req.Token); err != nil {
		return nil, err
	}
	userid, _ := tools.ParseToken(req.Token)
	commentList, err1 := s.GetList(req.VideoId, userid.User_id)
	if err1 != nil {
		msg = err1.Error()
		resp.BaseResp.StatusCode = 1
		resp.BaseResp.StatusMsg = &msg
		return resp, err1
	} else {
		msg = ""
		resp.BaseResp.StatusCode = 0
		resp.BaseResp.StatusMsg = &msg
		resp.CommentList = commentList
	}
	return resp, nil
}

func (s *CommentServiceImpl) GetList(videoId int64, userId int64) ([]*comment.Comment, error) {
	//1.先查询评论列表信息
	commentList, err := dao.GetCommentList(videoId)
	if err != nil {
		return nil, err
	}
	//当前有0条评论
	if commentList == nil {
		return nil, nil
	}

	//提前定义好切片长度
	commentInfoList := make([]*comment.Comment, len(commentList))

	wg := &sync.WaitGroup{}
	wg.Add(len(commentList))
	idx := 0
	for _, cmt := range commentList {
		//2.调用方法组装评论信息，再append
		var commentData comment.Comment
		//将评论信息进行组装，添加想要的信息,插入从数据库中查到的数据
		go func(comment dao.Comment) {
			oneComment(&commentData, &comment)
			commentInfoList[idx] = &commentData
			idx = idx + 1
			wg.Done()
		}(cmt)
	}
	wg.Wait()
	//评论排序-按照主键排序
	sort.Sort(CommentSlice(commentInfoList))

	return commentInfoList, nil
}

// 此函数用于给一个评论赋值：评论信息+用户信息 填充
func oneComment(cmt *comment.Comment, com *dao.Comment) {
	var wg sync.WaitGroup
	wg.Add(1)
	//根据评论用户id和当前用户id，查询评论用户信息
	var err error
	cmt.Id = com.Id
	cmt.Content = com.CommentText
	cmt.CreateDate = com.CreateDate.Format(constants.DateTime)
	userData, err := db_mysql.GetUserService().GetUserById(com.UserId)
	cmt.User = (*comment.User)(userData)

	if err != nil {
		log.Println("oneComment return err: " + err.Error()) //函数返回提示错误信息
	}
	wg.Done()
	wg.Wait()
}

// CommentSlice 此变量以及以下三个函数都是做排序-准备工作
type CommentSlice []*comment.Comment

func (a CommentSlice) Len() int { //重写Len()方法
	return len(a)
}
func (a CommentSlice) Swap(i, j int) { //重写Swap()方法
	a[i], a[j] = a[j], a[i]
}
func (a CommentSlice) Less(i, j int) bool { //重写Less()方法
	return a[i].Id > a[j].Id
}
