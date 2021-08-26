package ginuser

import (
	"Tranning_food/common"
	"Tranning_food/component"
	"Tranning_food/component/hasher"
	"Tranning_food/modules/user/userbiz"
	"Tranning_food/modules/user/usermodel"
	"Tranning_food/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := userstorage.NewSqlStorage(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))

	}
}
