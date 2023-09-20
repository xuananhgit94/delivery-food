package main

import (
	"food-delivery/component/appctx"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/upload/transport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	appContext := appctx.NewAppContext(db)
	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	r.Static("/static", "./static")
	// POST /restaurants
	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.UploadImage(appContext))

	restaurants := v1.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

	r.Run()
}
