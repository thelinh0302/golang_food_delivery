package restaurantbiz

import (
	"Tranning_food/common"
	"Tranning_food/modules/restaurant/restaurantmodel"
	"context"
	"fmt"
)

var idRestaurant int

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}
type LikeStoreGetRestaurant interface {
	GetRestaurantIdLike(ctx context.Context, id int) (map[int]int, error)
}

type getRestaurantBiz struct {
	store     GetRestaurantStore
	likeStore LikeStoreGetRestaurant
}

func NewGetRestaurantBiz(store GetRestaurantStore, likeStore LikeStoreGetRestaurant) *getRestaurantBiz {
	return &getRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context,
	id int,
) (*restaurantmodel.Restaurant, error) {

	result, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	idRestaurant = result.Id

	resLike, err := biz.likeStore.GetRestaurantIdLike(ctx, idRestaurant)

	if err != nil {
		fmt.Println("cannot get restaurant likes:", err)
	}

	if v := resLike; v != nil {
		fmt.Println(v)
		result.LikeCount = resLike[result.Id]
	}

	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	return result, nil
}
