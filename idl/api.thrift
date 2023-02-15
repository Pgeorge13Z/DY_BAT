namespace go api

struct User {
    1: required i64 id,             // 用户id
    2: required string name,        // 用户名称
    3: optional i64 follow_count,   // 关注总数
    4: optional i64 follower_count, // 粉丝总数
    5: required bool is_follow,     // true-已关注，false-未关注
}


// 基础返回信息 包括状态码和状态描述
struct BaseResp {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string statsu_msg, // 返回状态描述
}

// 用户注册接口
struct douyin_user_register_request {
    1: required string username (api.query="username",api.vd="len($) < 32"), // 注册用户名，最长32个字符
    2: required string password ( api.query="password",api.vd="len($) < 32" ), // 密码,最长32个字符
}

struct douyin_user_register_response {
    1: required i64 user_id,        // 用户id
    2: required string token,       // 用户鉴权token
    3: required BaseResp base_resp,
}

// 用户登陆接口
struct douyin_user_login_request {
    1: required string username ( api.query="username",api.vd="len($) < 32" ), // 登录用户名
    2: required string password ( api.query="username",api.vd="len($) < 32" ), // 登录密码
}

struct douyin_user_login_response {
    1: required i64 user_id,        // 用户id
    2: required string token,       // 用户鉴权token
    3: required BaseResp base_resp,
}

// 用户信息
struct DouyinUserRequest{
    1: required i64 user_id (api.query="user_id") //用户id
    2: required string token (api.query="token") //用户token
}

struct douyin_user_response {
    1: required User user,          // 用户信息
    2: required BaseResp base_resp,
}
