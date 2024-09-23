package registerbiz

import (
	"context"
	"social/common"
	usermodel "social/module/user/model"
)

type RegisterStorage interface {
	FindByUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) (*usermodel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	RegisterStorage RegisterStorage
	Hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		RegisterStorage: registerStorage,
		Hasher:          hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := business.RegisterStorage.FindByUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		if user.Status == 0 {
			return usermodel.ErrUserHasBeenDisabled
		}
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = business.Hasher.Hash(data.Password + salt)
	data.Salt = salt

	if user, err := business.RegisterStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCanNotCreateEntity(user.TableName(), err)
	}

	return nil
}
