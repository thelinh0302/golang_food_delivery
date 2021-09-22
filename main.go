package main

import (
	"Tranning_food/component"
	"Tranning_food/component/uploadprovider"
	"Tranning_food/middleware"
	"Tranning_food/modules/restaurant/restauranttransfort/ginrestaurant"
	"Tranning_food/modules/restaurantlikes/transport/ginrestaurantlike"
	"Tranning_food/modules/upload/uploadtransfort/ginupload"
	"Tranning_food/modules/user/usertransport/ginuser"
	"Tranning_food/pubsub/pblocal"
	"Tranning_food/subscriber"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	dsn := os.Getenv("DBConectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3ApiKey")
	s3SecretKey := os.Getenv("S3Secretkey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	fmt.Println(s3APIKey, s3SecretKey)

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	db = db.Debug()
	if err != nil {
		log.Fatal(err)
	}

	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {
	appCtx := component.NewAppContext(db, upProvider, secretKey, pblocal.NewPubSub())
	if err := subscriber.NewEngine(appCtx).Start(); err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
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
