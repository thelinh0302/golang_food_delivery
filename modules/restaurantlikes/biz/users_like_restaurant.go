package restaurantlikebiz

import (
	"Tranning_food/common"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	"context"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikesmodel.Like) error
	FindDataLikeByCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantlikesmodel.Like, error)
}

type IncreaseUserLikeRestaurantStore interface {
	IncreaseLike(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	increase IncreaseUserLikeRestaurantStore
}

func NewUserLikeRestaurantStore(store UserLikeRestaurantStore, increase IncreaseUserLikeRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, increase: increase}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikesmodel.Like) error {

	findData, err := biz.store.FindDataLikeByCondition(ctx, map[string]interface{}{"restaurant_id": data.RestaurantId, "user_id": data.UserId})

	if findData != nil {
		return common.ErrEntityExisted(restaurantlikesmodel.EntityName, err)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikesmodel.ErrCannotLikeRestaurant(err)
	}

	//side effect
	go func() {
		defer common.AppRecover()
		_ = biz.increase.IncreaseLike(ctx, data.RestaurantId)
	}()
	return nil
}
