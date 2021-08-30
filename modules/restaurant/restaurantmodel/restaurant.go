package restaurantmodel

import (
	"Tranning_food/common"
	"errors"
	"strings"
)

const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	UserId          int                `json:"-" gorm:"column:owner_id;"`
	Addr            string             `json:"address" gorm:"column:addr;"` //tag
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
	LikeCount       int                `json:"like_count" gorm:"-"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdated struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"address" gorm:"column:addr;"` //tag
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdated) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	UserId          int                `json:"-" gorm:"column:owner_id;"`
	Addr            string             `json:"address" gorm:"column:addr;"` //tag
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
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

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
	if u := data.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}
