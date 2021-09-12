package restaurantlikebiz

import (
	"Tranning_food/common"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	"context"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
	FindDataLikeByCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantlikesmodel.Like, error)
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
	_, err := biz.store.FindDataLikeByCondition(ctx, map[string]interface{}{"restaurant_id": restaurantId, "user_id": userId})

	if err != nil {
		return common.ErrEntityNotFound(restaurantlikesmodel.EntityName, err)
	}

	if err := biz.store.Delete(ctx, userId, restaurantId); err != nil {
		return restaurantlikesmodel.ErrCannotUnlikeRestaurant(err)

	}

	//side effect
	go func() {
		defer common.AppRecover()
		_ = biz.decrease.DecreaseLike(ctx, restaurantId)
	}()
	return nil
}
