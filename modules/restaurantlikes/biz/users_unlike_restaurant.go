package restaurantlikebiz

import (
	"Tranning_food/common"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	"context"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
}

type DecreaseUserUnLikeStore interface {
	DecreaseLike(ctx context.Context, id int) error
}

type userUnLikeRestaurantBiz struct {
	store    UserUnLikeRestaurantStore
	decrease DecreaseUserUnLikeStore
}

func NewUnUserLikeRestaurantStore(store UserUnLikeRestaurantStore, decrease DecreaseUserUnLikeStore) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{store: store, decrease: decrease}
}

func (biz *userUnLikeRestaurantBiz) UnLikeRestaurant(ctx context.Context, userId, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikesmodel.ErrCannotUnlikeRestaurant(err)
	}

	//side effect
	go func() {
		defer common.AppRecover()
		_ = biz.decrease.DecreaseLike(ctx, restaurantId)
	}()
	return nil
}
