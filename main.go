package main

import (
	"Tranning_food/component"
	otpServices2 "Tranning_food/component/otpProvider"
	"Tranning_food/component/uploadprovider"
	"Tranning_food/middleware"
	"Tranning_food/modules/restaurant/restauranttransfort/ginrestaurant"
	"Tranning_food/modules/restaurantlikes/transport/ginrestaurantlike"
	"Tranning_food/modules/upload/uploadtransfort/ginupload"
	"Tranning_food/modules/user/usertransport/ginuser"
	"Tranning_food/pubsub/pblocal"
	"Tranning_food/skio"
	"Tranning_food/subscriber"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := os.Getenv("DBConectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3ApiKey")
	s3SecretKey := os.Getenv("S3Secretkey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	accountID := os.Getenv("AccountSID")
	tokenID := os.Getenv("AuthToken")
	fromPhone := os.Getenv("FromPhone")
	otpProvider := otpServices2.NewOtpServicesProvider(accountID, tokenID, fromPhone)
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	db = db.Debug()
	if err != nil {
		log.Fatal(err)
	}

	if err := runService(db, s3Provider, secretKey, otpProvider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string, otpProvider otpServices2.OtpProvider) error {
	appCtx := component.NewAppContext(db, upProvider, secretKey, pblocal.NewPubSub(), otpProvider)
	r := gin.Default()

	rtEngine := skio.NewEngine()

	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Fatalln(err)
	}

	if err := subscriber.NewEngine(appCtx, rtEngine).Start(); err != nil {
		log.Fatalln(err)
	}

	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.StaticFile("/demo/", "./demo.html")

	//crud
	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appCtx))

	v1.POST("/register", ginuser.CreateUser(appCtx))

	v1.POST("/login", ginuser.Login(appCtx))

	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	{
		//POST
		restaurants.POST("", ginrestaurant.CreateResaurant(appCtx))
		//GET ALL HAVE FILTER OR NO FILTER
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		//UPDATE RESTAURANT
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		//DELETE
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
		//GET BY ID
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		//Get user like restaurant
		restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUser(appCtx))
		//Post v1/restaurants/:id/like
		restaurants.POST("/:id/like", ginrestaurantlike.UserLikeRestaurant(appCtx))
		//Delete v1/restaurants/:id/like
		restaurants.DELETE("/:id/unlike", ginrestaurantlike.UserUnLikeRestaurant(appCtx))

	}
	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//func startSocketIOServer(engine *gin.Engine, appCtx component.AppContext) {
//	server, _ := socketio.NewServer(&engineio.Options{
//		Transports: []transport.Transport{websocket.Default},
//	})
//
//	server.OnConnect("/", func(s socketio.Conn) error {
//		//s.SetContext("")
//		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())
//
//		//s.Join("Shipper")
//		//server.BroadcastToRoom("/", "Shipper", "test", "Hello 200lab")
//
//		return nil
//	})
//
//	server.OnError("/", func(s socketio.Conn, e error) {
//		fmt.Println("meet error:", e)
//	})
//
//	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
//		fmt.Println("closed", reason)
//		// Remove socket from socket engine (from app context)
//	})
//
//	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
//
//		// Validate token
//		// If false: s.Close(), and return
//
//		// If true
//		// => UserId
//		// Fetch db find user by Id
//		// Here: s belongs to who? (user_id)
//		// We need a map[user_id][]socketio.Conn
//
//		db := appCtx.GetMainDBConection()
//		store := userstorage.NewSqlStorage(db)
//		//
//		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
//		//
//		payload, err := tokenProvider.Validate(token)
//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//		//
//		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
//		//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//		//
//		if user.Status == 0 {
//			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
//			s.Close()
//			return
//		}
//		//
//		user.Mask(false)
//
//		s.Emit("your_profile", user)
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
//		log.Println(msg)
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
//
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
//		fmt.Println("server receive test:", msg)
//		s.Emit("test", msg)
//	})
//	//
//	//server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
//	//	s.SetContext(msg)
//	//	return "recv " + msg
//	//})
//	//
//	//server.OnEvent("/", "bye", func(s socketio.Conn) string {
//	//	last := s.Context().(string)
//	//	s.Emit("bye", last)
//	//	s.Close()
//	//	return last
//	//})
//	//
//	//server.OnEvent("/", "noteSumit", func(s socketio.Conn) string {
//	//	last := s.Context().(string)
//	//	s.Emit("bye", last)
//	//	s.Close()
//	//	return last
//	//})
//
//	go server.Serve()
//
//	engine.GET("/socket.io/*any", gin.WrapH(server))
//	engine.POST("/socket.io/*any", gin.WrapH(server))
//}
