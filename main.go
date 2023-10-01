package main

import (
	"food-delivery/component/appctx"
	"food-delivery/component/uploadprovider"
	"food-delivery/middleware"
	"food-delivery/pubsub/localpb"
	"food-delivery/subscriber"
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

	ps := localpb.NewPubSub()

	appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps)

	//subscriber.Setup(appContext, context.Background())
	_ = subscriber.NewEngine(appContext).Start()
	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	// POST /restaurants
	v1 := r.Group("/v1")

	setupRoute(appContext, v1, middleware.RequireAuth(appContext))
	setupAdmin(appContext, v1)

	r.Run()
}
