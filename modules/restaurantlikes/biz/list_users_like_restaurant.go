package restaurantlikebiz

import (
	"Tranning_food/common"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	"context"
)

type ListUserRestaurantLikeStore interface {
	GetUsersRestaurantLike(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantlikesmodel.Filter,
		paging *common.Paging,
		morekeys ...string) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantbBiz struct {
	store ListUserRestaurantLikeStore
}

func NewListUserLikeRestaurant(store ListUserRestaurantLikeStore) *listUserLikeRestaurantbBiz {
	return &listUserLikeRestaurantbBiz{store: store}
}

func (biz *listUserLikeRestaurantbBiz) ListUserLike(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikesmodel.Filter,
	paging *common.Paging,
	morekeys ...string) ([]common.SimpleUser, error) {

	users, err := biz.store.GetUsersRestaurantLike(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikesmodel.EntityName, err)
	}
	return users, nil

}
