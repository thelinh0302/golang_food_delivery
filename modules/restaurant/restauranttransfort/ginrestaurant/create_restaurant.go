package ginrestaurant

import (
	"Tranning_food/common"
	"Tranning_food/component"
	"Tranning_food/modules/restaurant/restaurantbiz"
	"Tranning_food/modules/restaurant/restaurantmodel"
	"Tranning_food/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateResaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSqlStorage(appCtx.GetMainDBConection())

		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
