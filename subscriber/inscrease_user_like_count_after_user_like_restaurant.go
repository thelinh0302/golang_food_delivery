package subscriber

import (
	"Tranning_food/component"
	"Tranning_food/modules/restaurant/restaurantstorage"
	"Tranning_food/pubsub"
	"Tranning_food/skio"
	"context"
)

type HasRestaunrantId interface {
	GetRestaurantId() int
	GetUserId() int
}

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

func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSqlStorage(appCtx.GetMainDBConection())
			likeData := message.Data().(HasRestaunrantId)
			return store.IncreaseLike(ctx, likeData.GetRestaurantId())
		},
	}
}
func EmitIncreaseLikeCountAfterUserLikeRestaurant(rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob{
		Title: "Emit realtime after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaunrantId)
			return rtEngine.EmitToUser(likeData.GetUserId(), string(message.Channel()), likeData)
		},
	}
}
