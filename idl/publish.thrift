include "user.thrift"

namespace go publish

struct Video {
    1: required i64 id,             // 视频唯一表示
    2: required user.User author,        // 视频作者信息
    3: required string play_url,    // 视频播放地址
    4: required string cover_url,   // 视频封面地址
    5: required i64 favorite_count, // 视频的点赞总数
    6: required i64 comment_count,  // 视频的评论总数
    7: required bool is_favorite,   // true-已点赞，false-未点赞
    8: required string title,       // 视频标题
}


struct douyin_publish_action_request  {
    1:required string token;
    2:required binary  data;
    3:required string title;
}

struct douyin_publish_action_response  {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string statsu_msg, // 返回状态描述
}

struct douyin_publish_list_request  {
       1: required i64 user_id ( vt.gt = "0" )
       2: required string token
}

struct douyin_publish_list_response   {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string statsu_msg, // 返回状态描述
    3: required list<Video> video_list,
}


service PublishService {
    douyin_publish_action_response PublishAction(1: douyin_publish_action_request req)
    douyin_publish_list_response PublishList(1: douyin_publish_list_request req)
}