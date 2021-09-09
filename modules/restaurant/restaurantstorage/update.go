package restaurantstorage

import (
	"Tranning_food/common"
	"Tranning_food/modules/restaurant/restaurantmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStorage) UpdateData(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdated,
) error {
	db := s.db

	if err := db.Where("id= ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStorage) IncreaseLike(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id= ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (s *sqlStorage) DecreaseLike(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id= ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
