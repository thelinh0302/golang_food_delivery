package restaurantrepo

import (
	"Tranning_food/common"
	"Tranning_food/modules/restaurant/restaurantmodel"
	"context"
	"fmt"
)

type ListRestaurantRepo interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeStore interface {
	GetRestaurantLike(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantRepo
	likeStore LikeStore
}

func NewListRestaurantRepo(store ListRestaurantRepo, likeStore LikeStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store:     store,
		likeStore: likeStore,
	}
}

func (biz *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}
	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := biz.likeStore.GetRestaurantLike(ctx, ids)

	if err != nil {
		fmt.Println("Cannot get restaurant likes:", err)
	}
	if v := mapResLike; v != nil {
		for i, item := range result {
			result[i].LikeCount = mapResLike[item.Id]
		}
	}

	return result, nil
}
