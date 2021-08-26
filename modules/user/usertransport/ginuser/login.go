package ginuser

import (
	"Tranning_food/common"
	"Tranning_food/component"
	"Tranning_food/component/hasher"
	"Tranning_food/component/tokenprovider/jwt"
	"Tranning_food/modules/user/userbiz"
	"Tranning_food/modules/user/usermodel"
	"Tranning_food/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSqlStorage(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, md5, tokenProvider, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
