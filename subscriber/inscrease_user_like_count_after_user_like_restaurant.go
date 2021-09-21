package subscriber

import (
	"Tranning_food/common"
	"Tranning_food/component"
	"Tranning_food/modules/restaurant/restaurantstorage"
	"context"
)

type HasRestaunrantId interface {
	GetRestaurantId() int
}

func IncreaseUserLikeCountAfterUserLikeRestaurant(appctx component.AppContext, ctx context.Context) {

	c, _ := appctx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)

	store := restaurantstorage.NewSqlStorage(appctx.GetMainDBConection())
	go func() {
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaunrantId)
			_ = store.IncreaseLike(ctx, likeData.GetRestaurantId())
		}
	}()

}
