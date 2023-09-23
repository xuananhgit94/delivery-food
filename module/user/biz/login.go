package userbiz

//
//import (
//	"context"
//	"food-delivery/common"
//	"food-delivery/component/appctx"
//	usermodel "food-delivery/module/user/model"
//)
//
//type LoginStorage interface {
//	FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
//}
//
//type loginBusiness struct {
//	appCtx        appctx.AppContext
//	storeUser     LoginStorage
//	tokenProvider tokenprovider.Provider
//	hasher        Hasher
//	expiry        int
//}
//
//func NewLoginBussiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
//	return &loginBusiness{
//		storeUser:     storeUser,
//		tokenProvider: tokenProvider,
//		hasher:        hasher,
//		expiry:        expiry,
//	}
//}
//
//func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
//	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
//	if err != nil {
//		panic(err)
//	}
//
//	passHashed := business.hasher.Hash(data.Password + user.Salt)
//
//	if user.Password != passHashed {
//		return nil, usermodel.ErrUsernameOrPasswordInvalid
//	}
//
//	paload := tokenprovider.TokenPayload{
//		UserId: user.Id,
//		Role:   user.Role,
//	}
//
//	acessToken, err := business.tokenProvider.Generate(paload, business.expiry)
//
//	if err != nil {
//		return nil, common.ErrInternal(err)
//	}
//
//}
