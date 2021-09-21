package ginrestaurantlike

import (
	"Tranning_food/common"
	"Tranning_food/component"
	restaurantlikebiz "Tranning_food/modules/restaurantlikes/biz"
	restaurantlikesmodel "Tranning_food/modules/restaurantlikes/model"
	restaurantlikestorage "Tranning_food/modules/restaurantlikes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Post v1/restaurants/:id/ike
func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data := restaurantlikesmodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}
		store := restaurantlikestorage.NewSqlStorage(appCtx.GetMainDBConection())

		biz := restaurantlikebiz.NewUserLikeRestaurantStore(store, appCtx.GetPubSub())

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
