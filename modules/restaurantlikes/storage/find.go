package restaurantlikestorage

import (
	"Tranning_food/common"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStorage) FindDataLikeByCondition(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*restaurantlikesmodel.Like, error) {

	var result restaurantlikesmodel.Like

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &result, nil
}
