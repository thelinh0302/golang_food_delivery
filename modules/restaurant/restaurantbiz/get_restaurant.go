package restaurantbiz

import (
	"Tranning_food/modules/restaurant/restaurantmodel"
	"context"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context,
	id int,
) (*restaurantmodel.Restaurant, error) {

	result, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}
	return result, nil
}
