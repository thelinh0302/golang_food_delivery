package restaurantstorage

import (
	"Tranning_food/common"
	"Tranning_food/modules/restaurant/restaurantmodel"
	"context"
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
