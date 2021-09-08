package restaurantrepo

import (
	"Tranning_food/common"
	"Tranning_food/modules/restaurant/restaurantmodel"
	"context"
	"fmt"
)

type GetRestaurantRepo interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type LikeStoreGetRestaurant interface {
	GetRestaurantIdLike(ctx context.Context, id int) (map[int]int, error)
}

type getRestaurantRepo struct {
	store     GetRestaurantRepo
	likeStore LikeStoreGetRestaurant
}

func NewGetRestaurantBiz(store GetRestaurantRepo, likeStore LikeStoreGetRestaurant) *getRestaurantRepo {
	return &getRestaurantRepo{store: store, likeStore: likeStore}
}

func (biz *getRestaurantRepo) GetRestaurant(ctx context.Context,
	id int,
) (*restaurantmodel.Restaurant, error) {

	result, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id}, "User")

	idRestaurant := result.Id

	resLike, err := biz.likeStore.GetRestaurantIdLike(ctx, idRestaurant)

	if err != nil {
		fmt.Println("cannot get restaurant likes:", err)
	}

	if v := resLike; v != nil {
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
