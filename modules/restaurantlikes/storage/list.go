package restaurantlikestorage

import (
	"Tranning_food/common"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	"context"
)

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
