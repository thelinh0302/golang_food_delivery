package userstorage

import (
	"Tranning_food/common"
	"Tranning_food/modules/user/usermodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStorage) FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var user usermodel.User
	if err := db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &user, nil
}
