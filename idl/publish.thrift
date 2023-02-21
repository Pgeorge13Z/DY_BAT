namespace go publish

struct Video {
    1: required i64 id,             // 视频唯一表示
    2: required User author,        // 视频作者信息
    3: required string play_url,    // 视频播放地址
    4: required string cover_url,   // 视频封面地址
    5: required i64 favorite_count, // 视频的点赞总数
    6: required i64 comment_count,  // 视频的评论总数
    7: required bool is_favorite,   // true-已点赞，false-未点赞
    8: required string title,       // 视频标题
}

struct User {
    1: required i64 id,             // 用户id
    2: required string name,        // 用户名称
    3: optional i64 follow_count,   // 关注总数
    4: optional i64 follower_count, // 粉丝总数
    5: required bool is_follow,     // true-已关注，false-未关注
    6: optional string avatar, //用户头像
    7: optional string background_image,//用户个人页顶部大图
    8: optional string signature, //个人简介
    9: optional i64 totaol_favorited, //获赞数量
    10: optional i64 work_count, //作品数量
    11: optional i64 favorite_count, //点赞数量
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