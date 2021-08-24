package userbiz

import (
	"Tranning_food/common"
	"Tranning_food/modules/user/usermodel"
	"context"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}
type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{registerStorage: registerStorage, hasher: hasher}
}

func (register *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {

	user, err := register.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return common.ErrEntityExisted(usermodel.EntityName, err)
	}

	salt := common.GenSalt(50)
	data.Password = register.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1
	if err := register.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
