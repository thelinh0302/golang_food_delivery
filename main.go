package main

import (
	"Tranning_food/component"
	"Tranning_food/modules/restaurant/restauranttransfort/ginrestaurant"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"` //tag
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdated struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"` //tag
}

func (RestaurantUpdated) TableName() string {
	return Restaurant{}.TableName()
}

func main() {

	dsn := os.Getenv("DBConectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)

	if err != nil {
		log.Fatal(err)
	}

	if err := runService(db); err != nil {
		log.Fatal(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//crud

	appCtx := component.NewAppContext(db)
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
		restaurants.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(401, gin.H{ //H.ginrestaurant{}
					"error": err.Error(),
				})
				return
			}
			var data Restaurant
			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, data)
		})

	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
