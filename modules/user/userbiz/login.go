package userbiz

import (
	"Tranning_food/common"
	"Tranning_food/component"
	"Tranning_food/component/tokenprovider"
	"Tranning_food/modules/user/usermodel"
	"context"
)

type LoginStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type TokenConfig interface {
	GetAtExp() int
	GetRtExp() int
}

type logiBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	hasher        Hasher
	expiry        int
	tokenProvider tokenprovider.Provider
}

func NewLoginBusiness(
	storeUser LoginStorage,
	hasher Hasher,
	tokenProvider tokenprovider.Provider,
	expiry int) *logiBusiness {
	return &logiBusiness{
		storeUser:     storeUser,
		hasher:        hasher,
		tokenProvider: tokenProvider,
		expiry:        expiry,
	}
}

//1.Find user,email
//2.Hash pass from input and compare pass in db
//3.Provider: issue JWT Token for client
//3.1 Access token and refresh token
//4 Return token

func (business *logiBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	passHasher := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHasher {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}
	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	//refresToken, err := business.tokenProvider.Generate(payload, business.tkCfg.GetRtExp())
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}
	//account := usermodel.NewAccount(accessToken, refresToken)
	return accessToken, nil
}
