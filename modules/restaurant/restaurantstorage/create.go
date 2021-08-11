package restaurantstorage

import (
	"Tranning_food/common"
	"Tranning_food/modules/restaurant/restaurantmodel"
	"context"
)

func (s *sqlStorage) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
