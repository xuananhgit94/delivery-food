package main

import (
	"food-delivery/component/appctx"
	"food-delivery/component/uploadprovider"
	"food-delivery/middleware"
	"food-delivery/pubsub/localpb"
	"food-delivery/skio"
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

	r.StaticFile("/demo/", "./demo.html")

	r.Use(middleware.Recover(appContext))

	// POST /restaurants
	v1 := r.Group("/v1")

	setupRoute(appContext, v1, middleware.RequireAuth(appContext))
	setupAdmin(appContext, v1)

	//startSocketIOServer(r, appContext)

	rtEngine := skio.NewEngine()
	appContext.SetRealtimeEngine(rtEngine)

	_ = rtEngine.Run(appContext, r)

	r.Run()
}

//func startSocketIOServer(engine *gin.Engine, appCtx appctx.AppContext) {
//	server, _ := socketio.NewServer(&engineio.Options{
//		Transports: []transport.Transport{websocket.Default},
//	})
//
//	server.OnConnect("/", func(s socketio.Conn) error {
//		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())
//
//		s.Join("Shipper")
//
//		return nil
//	})
//
//	//go func() {
//	//	for range time.NewTicker(time.Second).C {
//	//		server.BroadcastToRoom("/", "Shipper", "test", "Ahihi")
//	//	}
//	//}()
//
//	server.OnError("/", func(s socketio.Conn, e error) {
//		fmt.Println("meet error:", e)
//	})
//
//	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
//		fmt.Println("closed", reason)
//	})
//
//	server.OnEvent("/", "authenticated", func(s socketio.Conn, token string) {
//		db := appCtx.GetMainDBConnection()
//		store := userstore.NewSQLStore(db)
//
//		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
//		payload, err := tokenProvider.Validate(token)
//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//
//		if user.Status == 0 {
//			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
//			s.Close()
//			return
//		}
//		user.Mask(false)
//
//		s.Emit("your_profile", user)
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg interface{}) {
//		log.Println("test:", msg)
//	})
//
//	type Person struct {
//		Name string `json:"name"`
//		Age  int    `json:"age"`
//	}
//
//	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) {
//		fmt.Println("server receive notice:", p.Name, p.Age)
//
//		p.Age = 33
//		s.Emit("notice", p)
//	})
//
//	go server.Serve()
//
//	engine.GET("/socket.io/*any", gin.WrapH(server))
//	engine.POST("/socket.io/*any", gin.WrapH(server))
//}
