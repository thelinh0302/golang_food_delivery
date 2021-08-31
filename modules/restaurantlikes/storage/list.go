package restaurantlikestorage

import (
	"Tranning_food/common"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	"context"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

type sqlData struct {
	RestaurantId int `gorm:"column:restaurant_id;"`
	LikeCount    int `gorm:"column:count;"`
}

func (s *sqlStorage) GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	var likeCount []sqlData

	if err := s.db.Table(restaurantlikesmodel.Like{}.TableName()).
		Select("restaurant_id,count(restaurant_id) as count").
		Group("restaurant_id").
		Where("restaurant_id in (?)", ids).Find(&likeCount).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range likeCount {
		result[item.RestaurantId] = item.LikeCount
	}
	return result, nil
}

func (s *sqlStorage) GetRestaurantIdLike(ctx context.Context, id int) (map[int]int, error) {
	result := make(map[int]int)

	var likeCount []sqlData

	if err := s.db.Table(restaurantlikesmodel.Like{}.TableName()).
		Select("restaurant_id,count(restaurant_id) as count").
		Group("restaurant_id").
		Where("restaurant_id", id).Find(&likeCount).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	for _, item := range likeCount {
		result[item.RestaurantId] = item.LikeCount
	}
	return result, nil
}

func (s *sqlStorage) GetUsersRestaurantLike(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikesmodel.Filter,
	paging *common.Paging,
	morekeys ...string) ([]common.SimpleUser, error) {

	var result []restaurantlikesmodel.Like
	db := s.db

	db = db.Table(restaurantlikesmodel.Like{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id =?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User
		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}
