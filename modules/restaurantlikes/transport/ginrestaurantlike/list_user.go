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

func ListUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("id"))
		//var filter restaurantlikesmodel.Filter
		//
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikesmodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSqlStorage(appCtx.GetMainDBConection())

		biz := restaurantlikebiz.NewListUserLikeRestaurant(store)

		users, err := biz.ListUserLike(c.Request.Context(), nil, &filter, &paging)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range users {
			users[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, filter))

	}

}
