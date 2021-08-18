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

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.Fulfill()
		store := restaurantstorage.NewSqlStorage(appCtx.GetMainDBConection())
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
