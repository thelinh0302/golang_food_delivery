package ginrestaurantlike

import (
	"Tranning_food/common"
	"Tranning_food/component"
	restaurantlikebiz "Tranning_food/modules/restaurantlikes/biz"
	restaurantlikestorage "Tranning_food/modules/restaurantlikes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Delete v1/restaurants/:id/unlike
func UserUnLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSqlStorage(appCtx.GetMainDBConection())

		biz := restaurantlikebiz.NewUnUserLikeRestaurantStore(store, appCtx.GetPubSub())

		if err := biz.UnLikeRestaurant(c.Request.Context(),
			requester.GetUserId(),
			int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
