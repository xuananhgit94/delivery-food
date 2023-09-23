package main

import (
	"food-delivery/component/appctx"
	"food-delivery/component/uploadprovider"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/upload/transport/ginupload"
	"food-delivery/module/user/transport/ginuser"
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

	s3BucketName := os.Getenv("S3BucketName")
	s3Endpoint := os.Getenv("S3Endpoint")
	s3AccessKeyID := os.Getenv("S3AccessKeyID")
	s3SecretAccessKey := os.Getenv("S3secretAccessKey")
	s3Region := os.Getenv("S3Region")

	secretKey := os.Getenv("SYSTEM_SECRET")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Endpoint, s3AccessKeyID, s3SecretAccessKey, s3Region)

	appContext := appctx.NewAppContext(db, s3Provider, secretKey)
	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	// POST /restaurants
	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appContext))

	v1.POST("/register", ginuser.Register(appContext))

	v1.POST("/authenticate", ginuser.Login(appContext))

	v1.GET("/profile", middleware.RequireAuth(appContext), ginuser.Profile(appContext))

	restaurants := v1.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

	r.Run()
}
