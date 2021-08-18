package ginrestaurant

import (
	"Tranning_food/common"
	"Tranning_food/component"
	"Tranning_food/modules/restaurant/restaurantbiz"
	"Tranning_food/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))
		//id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSqlStorage(appCtx.GetMainDBConection())

		biz := restaurantbiz.NewGetRestaurantBiz(store)
		result, err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}
		result.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
