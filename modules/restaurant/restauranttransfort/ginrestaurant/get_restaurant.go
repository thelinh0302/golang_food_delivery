package ginrestaurant

import (
	"Tranning_food/common"
	"Tranning_food/component"
	"Tranning_food/modules/restaurant/restaurantbiz"
	"Tranning_food/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		store := restaurantstorage.NewSqlStorage(appCtx.GetMainDBConection())

		biz := restaurantbiz.NewGetRestaurantBiz(store)
		result, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}