package db

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

// select the video id that the user likes
func SelectLikeVideoIds(userId int64) ([]int64, error) {
	var ids []int64
	if err := db.Model(Favorites{}).
		Where(" user_id= ? and is_favorite = ?", userId, 1).
		Pluck("video_id", &ids).Error; nil != err {
		if "record not found" == err.Error() {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, err
	}
	return ids, nil
}

func SelectFavoriteRecord(userID int64, videoID int64) (*Favorites, error) {
	var likesRecord Favorites
	if err := db.Model(&Favorites{}).
		Where("user_id = ? and video_id = ?", userID, videoID).
		First(&likesRecord).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return &likesRecord, err
	}
	return &likesRecord, nil
}

func InsertFavoriteRecord(userID int64, videoID int64, action_type int32) (bool, error) {
	likesRecord := Favorites{
		UserId:     userID,
		VideoId:    videoID,
		IsFavorite: action_type,
	}
	if err := db.Create(&likesRecord).Error; err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}

func UpdateFavoriteRecord(userID int64, videoID int64, action_type int32) (bool, error) {
	if err := db.Model(Favorites{}).
		Where("user_id = ? and video_id = ?", userID, videoID).
		Update("is_favorite", action_type).Error; nil != err {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}

func SelectVideobyId(videoID int64) (Videos, error) {
	v := Videos{}
	/*db.Where("name <> ?", "jinzhu").Find(&users)*/
	if err := db.Model(Videos{}).
		Where(" video_id= ?", videoID).
		Find(&v).Error; nil != err {
		if "record not found" == err.Error() {
			return v, nil
		}
		log.Println(err.Error())
		return v, err
	}
	return v, nil
}
