package restaurantbiz

import (
	"Tranning_food/common"
	"Tranning_food/modules/restaurant/restaurantmodel"
	"context"
)

type GetRestaurantRepo interface {
	GetRestaurant(ctx context.Context,
		id int,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	repo GetRestaurantRepo
}

func NewGetRestaurantBiz(repo GetRestaurantRepo) *getRestaurantBiz {
	return &getRestaurantBiz{repo: repo}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context,
	id int,
) (*restaurantmodel.Restaurant, error) {

	result, err := biz.repo.GetRestaurant(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
