package ginuser

import (
	"Tranning_food/common"
	"Tranning_food/component"
	"Tranning_food/component/tokenprovider/jwt"
	"Tranning_food/modules/user/userbiz"
	"Tranning_food/modules/user/usermodel"
	"Tranning_food/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOtp(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var params usermodel.UserOTP

		if err := c.ShouldBind(&params); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := ctx.GetMainDBConection()
		tokenProvider := jwt.NewTokenJWTProvider(ctx.SecretKey())

		//providerSendMsg := messprovider.NewSendMessage("AC12d527b67b01ce3ff91a1932bba0b1e6", "e7be1f6f054c67060b239b4fdbbace63", "+18557971273")

		store := userstorage.NewSqlStorage(db)

		biz := userbiz.NewGetOtpBiz(store, tokenProvider, 60*60*24)
		account, err := biz.GetOtp(c.Request.Context(), &params)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
