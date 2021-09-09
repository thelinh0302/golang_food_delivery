package restaurantlikestorage

import (
	"Tranning_food/common"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	"context"
)

func (s *sqlStorage) Delete(ctx context.Context, userId, restaurantId int) error {
	db := s.db

	if err := db.Table(restaurantlikesmodel.Like{}.TableName()).
		Where("user_id= ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
