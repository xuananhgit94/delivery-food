package main

import (
	"food-delivery/component/appctx"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/restaurantlike/transport/ginrstlike"
	"food-delivery/module/upload/transport/ginupload"
	"food-delivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func setupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup, auth func(c *gin.Context)) {
	v1.POST("/upload", ginupload.Upload(appContext))

	v1.POST("/register", ginuser.Register(appContext))

	v1.POST("/authenticate", ginuser.Login(appContext))

	v1.GET("/profile", ginuser.Profile(appContext))

	restaurants := v1.Group("/restaurants", auth)
	{

		restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

		restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

		restaurants.POST("/:id/like", ginrstlike.UserLikeRestaurant(appContext))

		restaurants.DELETE("/:id/liked-users", ginrstlike.UserDislikeRestaurant(appContext))

		restaurants.GET("/:id/liked-users", ginrstlike.ListUser(appContext))
	}

}
