package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	"food-delivery/component/hasher"
	"food-delivery/component/tokenprovider/jwt"
	userbiz "food-delivery/module/user/biz"
	usermodel "food-delivery/module/user/model"
	userstore "food-delivery/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin
		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		bussiness := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := bussiness.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
