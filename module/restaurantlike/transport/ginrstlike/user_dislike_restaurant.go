package ginrstlike

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantstorage "food-delivery/module/restaurant/storage"
	rstlikebiz "food-delivery/module/restaurantlike/biz"
	restaurantlikestorage "food-delivery/module/restaurantlike/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DELETE /v1/restaurants/:id/dislike

func UserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())

		decStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := rstlikebiz.NewUserDislikeRestaurantBiz(store, decStore)

		if err := biz.DislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
