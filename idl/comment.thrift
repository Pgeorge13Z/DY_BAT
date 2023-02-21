namespace go comment

struct Comment {
    1: required i64 id,            // 视频评论id
    2: required User user,        // 评论用户信息
    3: required string content,   // 评论内容
    4: required string create_date, // 评论发布日期，格式 mm-dd
}

struct User {
    1: required i64 id, // 用户id
    2: required string name, // 用户名称
    3: optional i64 follow_count, // 关注总数
    4: optional i64 follower_count, // 粉丝总数
    5: required bool is_follow, // true-已关注，false-未关注
}

// 基础返回信息 包括状态码和状态描述
struct BaseResp {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string status_msg, // 返回状态描述
}

// 评论操作
struct douyin_comment_action_request {
    1: required string token,   // 用户鉴权token
    2: required i64 video_id, // 视频id
    3: required i32 action_type, // 1-发布评论，2-删除评论
    4: optional string comment_text, // 用户填写的评论内容，在action_type=1的时候使用
    5: optional i64 comment_id, // 要删除的评论id，在action_type=2的时候使用
}

struct douyin_comment_action_response {
    1: required BaseResp base_resp,
    2: optional Comment comment, // 评论成功返回评论内容，不需要重新拉取整个列表
}

// 视频评论列表
struct douyin_comment_list_request {
    1: required string token, // 用户鉴权token
    2: required i64 video_id, // 视频id
}

struct douyin_comment_list_response {
    1: required BaseResp base_resp,
    2: required list<Comment> comment_list, // 评论列表
}

service CommentService {
    douyin_comment_action_response CommentAction(1: douyin_comment_action_request req),
    douyin_comment_list_response CommentList(1: douyin_comment_list_request req),
}
