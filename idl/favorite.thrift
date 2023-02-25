namespace go favorite
include "feed.thrift"

struct douyin_favorite_action_request{
  1: required string token
  2: required i64 video_id
  3: required i32 action_type
}

struct douyin_favorite_action_response{
  1: required i32 status_code
  2: optional string status_msg
}

struct douyin_favorite_list_request{
  1: required i64 user_id
  2: required string token
}

struct douyin_favorite_list_response{
  1: required i32 status_code
  2: string status_msg
  3: list<feed.Video> video_list
}


service FavoriteService{
  douyin_favorite_action_response FavoriteAction(1: douyin_favorite_action_request req)
  douyin_favorite_list_response FavoriteList(1: douyin_favorite_list_request req)

}