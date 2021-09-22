package restaurantlikebiz

import (
	"Tranning_food/common"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	"Tranning_food/pubsub"
	"context"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restaurantId int) error
	FindDataLikeByCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantlikesmodel.Like, error)
}

//type DecreaseUserUnLikeStore interface {
//	DecreaseLike(ctx context.Context, id int) error
//}

type userUnLikeRestaurantBiz struct {
	store UserUnLikeRestaurantStore
	pb    pubsub.PubSub
	//decrease DecreaseUserUnLikeStore
}

func NewUnUserLikeRestaurantStore(
	store UserUnLikeRestaurantStore,
	pb pubsub.PubSub,
	//decrease DecreaseUserUnLikeStore,
) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{
		store: store,
		pb:    pb,
		//decrease: decrease,
	}
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
	//go func() {
	//	defer common.AppRecover()
	//	job := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.decrease.DecreaseLike(ctx, restaurantId)
	//	})
	//	_ = asyncjob.NewGroup(true, job).Run(ctx)
	//}()

	biz.pb.Publish(ctx, common.TopicUserDisLikeRestaurant, pubsub.NewMessage(&restaurantlikesmodel.Like{
		RestaurantId: restaurantId,
		UserId:       userId,
	}))

	return nil
}
