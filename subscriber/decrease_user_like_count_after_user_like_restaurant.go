package subscriber

import (
	"Tranning_food/component"
	"Tranning_food/modules/restaurant/restaurantstorage"
	"Tranning_food/pubsub"
	"context"
)

//func IncreaseUserLikeCountAfterUserLikeRestaurant(appctx component.AppContext, ctx context.Context) {
//
//	c, _ := appctx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
//
//	store := restaurantstorage.NewSqlStorage(appctx.GetMainDBConection())
//	go func() {
//		for {
//			msg := <-c
//			likeData := msg.Data().(HasRestaunrantId)
//			_ = store.IncreaseLike(ctx, likeData.GetRestaurantId())
//		}
//	}()
//}

func RunDescreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Descrease like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStorage(appCtx.GetMainDBConection())
			likeData := message.Data().(HasRestaunrantId)
			return store.DecreaseLike(ctx, likeData.GetRestaurantId())
		},
	}
}
