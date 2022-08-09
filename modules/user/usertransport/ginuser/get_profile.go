package ginuser

import (
	"Tranning_food/common"
	"Tranning_food/component"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}