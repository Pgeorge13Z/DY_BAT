package dao

import "log"

// 查询表中是否有user关注过target的记录并返回
func FindRelation(userId int64, targetId int64) (*Follow, error) {
	follow := Follow{}
	if err := db.
		Where("user_id = ? and follower_id = ?", "userId", "targetId").
		First(&follow).Error; nil != err {
		if "record not found" == err.Error() {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, err
	}
	return &follow, nil
}

func IfFollow(userId int64, targetId int64) (bool, error) {
	follow, err := FindRelation(userId, targetId)
	if err != nil {
		return false, err
	}
	if follow == nil || follow.Cancel != 0 {
		return false, nil
	}
	return true, nil
}

// 插入user关注target的数据
func InserFollowRelation(userId int64, targetId int64) (bool, error) {
	follow := Follow{
		UserId:     userId,
		FollowerId: targetId,
		Cancel:     0,
	}
	if err := db.Create(&follow).Error; err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}

// 更改关注关系，0为关注，1为取消
func UpdateFollowRelation(userid int64, targetId int64, cancel int8) (bool, error) {
	if err := db.Model(Follow{}).
		Where("user_id = ? and follower_id = ?", "userId", "targetId").
		Update("cancel", "cancel").Error; nil != err {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}

// 查询用户关注了哪些id
func FindFollowIds(followerId int64) ([]int64, error) {
	var ids []int64
	if err := db.Model(Follow{}).
		Where("follower_id = ? and cancel = ?", "followerId", "0").
		Pluck("user_id", &ids).Error; nil != err {
		if "record not found" == err.Error() {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, err
	}
	return ids, nil
}

// 查询哪些id关注了该用户id
func FindFollowerIds(userId int64) ([]int64, error) {
	var ids []int64
	if err := db.Model(Follow{}).
		Where("user_id = ? and cancel = ?", "userId", "0").
		Pluck("follower_id", &ids).Error; nil != err {
		if "record not found" == err.Error() {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, err
	}
	return ids, nil
}

// 查询userId关注数
func GetFollowCount(followerId int64) (int64, error) {
	var count int64
	if err := db.Model(Follow{}).
		Where("follower_id = ?, cancel = ?", "followerId", "0").
		Count(&count).Error; nil != err {
		log.Println(err.Error())
		return 0, err
	}
	return count, nil
}

// 查询userId粉丝数
func GetFollowerCount(userId int64) (int64, error) {
	var count int64
	if err := db.Model(Follow{}).
		Where("user_id = ?, cancel = ?", "userId", "0").
		Count(&count).Error; nil != err {
		log.Println(err.Error())
		return 0, err
	}
	return count, nil
}
