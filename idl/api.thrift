namespace go api

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

struct Video_User {

}

struct Video {
    1: required i64 id,             // 视频唯一表示
    //2: required User author,        // 视频作者信息
    2: required i64 userid
    3: required string play_url,    // 视频播放地址
    4: required string cover_url,   // 视频封面地址
    5: required i64 favorite_count, // 视频的点赞总数
    6: required i64 comment_count,  // 视频的评论总数
    7: required bool is_favorite,   // true-已点赞，false-未点赞
    8: required string title,       // 视频标题
}




// 用户注册接口
struct douyin_user_register_request {
    1: required string username (api.query="username",api.vd="len($) < 32"), // 注册用户名，最长32个字符
    2: required string password ( api.query="password",api.vd="len($) < 32" ), // 密码,最长32个字符
}

struct douyin_user_register_response {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required i64 user_id,        // 用户id
    4: required string token,       // 用户鉴权token
}

// 用户登陆接口
struct douyin_user_login_request {
    1: required string username ( api.query="username",api.vd="len($) < 32" ), // 登录用户名
    2: required string password ( api.query="username",api.vd="len($) < 32" ), // 登录密码
}

struct douyin_user_login_response {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required i64 user_id,        // 用户id
    4: required string token,       // 用户鉴权token
}

// 用户信息
struct Douyin_User_Request{
    1: required i64 user_id (api.query="user_id"), //用户id
    2: required string token (api.query="token"), //用户token
}

struct douyin_user_response {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required User user,          // 用户信息
}

// 视频投稿
struct douyin_publish_action_request {
    1: required string token (api.form="token"),
    2: required binary data (api.form="data"),                        // 视频数据

    3: required string title ( api.form="title", api.vd="len($) > 0" ), // 视频标题
}

struct douyin_publish_action_response {
        1: required i32 status_code,
        2: optional string status_msg,
}

// 发布列表
struct douyin_publish_list_request {
    1: required i64 user_id (api.query="user_id"), // 用户id
    2: required string token (api.query="token"),
}

struct douyin_publish_list_response {
    1: required list<Video> video_list, // 用户发布的视频列表
    2: required i32 status_code,
    3: optional string status_msg,
}

service ApiService {
    // user
    douyin_user_register_response UserRegister(1: douyin_user_register_request req) (api.post="/douyin/user/register/")
    douyin_user_login_response UserLogin(1: douyin_user_login_request req) (api.post="/douyin/user/login/")
    douyin_user_response UserInfo(1: Douyin_User_Request req) (api.get="/douyin/user/")

    // publish
    douyin_publish_action_response PublishAction(1: douyin_publish_action_request req) (api.post="/douyin/publish/action/")
    douyin_publish_list_response PublishList(1: douyin_publish_list_request req) (api.get="/douyin/publish/list/")
}