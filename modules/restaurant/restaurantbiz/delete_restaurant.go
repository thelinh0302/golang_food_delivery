package restaurantbiz

import (
	"Tranning_food/modules/restaurant/restaurantmodel"
	"context"
	"errors"
)

type DeleteRestaurantStore interface {
	DeleteData(ctx context.Context, id int) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("Data deleted")
	}
	if err := biz.store.DeleteData(ctx, id); err != nil {
		return err
	}
	return nil
}
