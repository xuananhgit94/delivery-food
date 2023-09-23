package userbiz

import (
	"context"
	"food-delivery/common"
	usermodel "food-delivery/module/user/model"
)

type RegisterStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStore RegisterStore
	hasher        Hasher
}

func NewRegisterBusiness(registerStorage RegisterStore, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStore: registerStorage,
		hasher:        hasher,
	}
}

func (bussiness *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := bussiness.registerStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = bussiness.hasher.Hash(data.Password + salt)

	data.Salt = salt
	data.Role = "user"
	//data.Status = 1

	if err := bussiness.registerStore.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
