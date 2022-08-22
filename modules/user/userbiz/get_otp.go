package userbiz

import (
	"Tranning_food/component"
	"Tranning_food/component/digitProvider"
	"Tranning_food/component/tokenprovider"
	"Tranning_food/modules/user/usermodel"
	"context"
)

type GetOtpStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type getOtpBiz struct {
	appCtx        component.AppContext
	storeUser     GetOtpStorage
	tokenProvider tokenprovider.Provider
	expiry        int
}

func NewGetOtpBiz(storeUser GetOtpStorage, tokenProvider tokenprovider.Provider, expired int) *getOtpBiz {
	return &getOtpBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		expiry:        expired,
	}
}

// 1. check phone exited
// 2. random otp
func (biz *getOtpBiz) GetOtp(ctx context.Context, data *usermodel.UserOTP) (*usermodel.ResOTP, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"phone": data.Phone})
	println(user)
	if err := data.ValidateOTP(); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, usermodel.ErrPhoneInvalid
	}

	otp, _ := digitProvider.GenerateOT1P()
	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}
	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	dataRes := &usermodel.ResOTP{
		AccessToken: accessToken,
		Otp:         otp,
	}

	println(otp)
	println(accessToken)

	return dataRes, nil
}
