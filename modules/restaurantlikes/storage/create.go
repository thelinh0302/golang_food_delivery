package restaurantlikestorage

import (
	"Tranning_food/common"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	"context"
)

func (s *sqlStorage) Create(ctx context.Context, data *restaurantlikesmodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
