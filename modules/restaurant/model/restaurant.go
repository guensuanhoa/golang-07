package restaurantmodel

import (
	"errors"

	"example.com/g07-food-delivery/common"
)

type Restaurant struct {
	common.SQLModel        /// Day goi la embed
	Name            string `json:"name" gorm:"column:name;"`
	Address         string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

var (
	ErrorNameCannotBeBlank    = errors.New("restaurant name cannot be blank")
	ErrorAddressCannotBeBlank = errors.New("restaurant address cannot be blank")
)
