namespace go follow

struct User {
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
}
        
struct douyin_follow_action_request {
    1:string token //用户鉴权token
    2:i64 to_user_id //对方用户id
    3:string action_type //1-关注，2-取消关注
}

struct douyin_follow_action_response {
    1:i32 status_code  //状态码 0成功，其他值失败
    2:string status_msg //返回状态描述
}

struct douyin_follow_list_request {
    1:i64 user_id  //用户id 
    2:string token //用户鉴权token
}

struct douyin_follow_list_response {
    1:i32 status_code //状态码，0-成功，其他值-失败
    2:string status_msg //返回状态描述
    3:list<User> user_list //用户信息列表
}

service FollowService {
    douyin_follow_action_response FollowAction(1:douyin_follow_action_request req)
    douyin_follow_list_response FollowList(1:douyin_follow_list_request req)
}