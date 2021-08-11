package restaurantmodel

import (
	"Tranning_food/common"
	"errors"
	"strings"
)

const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"address" gorm:"column:addr;"` //tag
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

type RestaurantCreate struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"` //tag
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("Restaurant cannot be blank")
	}
	return nil
}
