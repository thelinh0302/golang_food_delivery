package subscriber

import (
	"Tranning_food/component"
	"context"
)

func Setup(ctx component.AppContext) {
	IncreaseUserLikeCountAfterUserLikeRestaurant(ctx, context.Background())

}
