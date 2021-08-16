package main

import (
	"Tranning_food/component"
	"Tranning_food/component/uploadprovider"
	"Tranning_food/middleware"
	"Tranning_food/modules/restaurant/restauranttransfort/ginrestaurant"
	"Tranning_food/modules/upload/uploadtransfort/ginupload"
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

	fmt.Println(s3APIKey, s3SecretKey)

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)

	if err != nil {
		log.Fatal(err)
	}

	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, upProvider)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//crud
	r.POST("/upload", ginupload.Upload(appCtx))

	restaurants := r.Group("/restaurants")
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

	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
