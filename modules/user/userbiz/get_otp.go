package userbiz

import (
	"Tranning_food/component"
	"Tranning_food/component/messprovider"
	"Tranning_food/component/tokenprovider"
	"Tranning_food/modules/user/usermodel"
	"context"
)

type GetOtpStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type getOtpBiz struct {
	appCtx         component.AppContext
	storeUser      GetOtpStorage
	tokenProvider  tokenprovider.Provider
	expiry         int
	sendMsgProvder messprovider.MessProvider
}

func NewGetOtpBiz(storeUser GetOtpStorage, tokenProvider tokenprovider.Provider, expired int) *getOtpBiz {
	return &getOtpBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		expiry:        expired,
		//sendMsgProvder: sendMsgProvder,
	}
}

// 1. check phone exited
// 2. random otp
func (biz *getOtpBiz) GetOtp(ctx context.Context, data *usermodel.UserOTP) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"phone": data.Phone})
	println(user)
	if err := data.ValidateOTP(); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, usermodel.ErrPhoneInvalid
	}

	//otp, _ := digitProvider.GenerateOT1P()
	sendMSg, errOtp := biz.sendMsgProvder.SendMessage(ctx, "0399172329", "23432")
	if errOtp != nil {
		return nil, errOtp
	}
	println(sendMSg)
	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	return accessToken, nil
}
